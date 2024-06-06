package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
)

type Client struct {
	Nickname string
	conn     *websocket.Conn
	ctx      context.Context
}

type Message struct {
	From    string
	Content string
	SentAt  string
}

var (
	clients     map[*Client]bool = make(map[*Client]bool)
	joinCh      chan *Client     = make(chan *Client)
	broadcastCh chan Message     = make(chan Message)
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	nickname := r.URL.Query().Get("nickname")

	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Println(err)
		return
	}

	go broadcast()
	go joiner() //criando uma nova goroutine

	//um cliente conectou com o websocket
	newClient := Client{Nickname: nickname, conn: conn, ctx: r.Context()}
	joinCh <- &newClient

	reader(&newClient)
}

func reader(newClient *Client) {
	for {
		_, data, err := newClient.conn.Read(newClient.ctx)
		if err != nil {
			log.Println("Encerrando conexão do cliente " + err.Error())
			delete(clients, newClient)
			broadcastCh <- Message{From: newClient.Nickname, Content: newClient.Nickname + " saiu do chat", SentAt: time.Now().Format("02-01-2006 15:04:05")}
			break
		}

		var msgRec Message
		err = json.Unmarshal(data, &msgRec)

		if err != nil {
			log.Println("Erro ao fazer o unmarshalling da resposta " + err.Error())
			return
		}

		broadcastCh <- Message{From: msgRec.From, Content: msgRec.Content, SentAt: time.Now().Format("02-01-2006 15:04:05")}
	}
}

func joiner() {
	for newClient := range joinCh {
		clients[newClient] = true

		broadcastCh <- Message{From: newClient.Nickname, Content: "O usuário " + newClient.Nickname + " se conectou", SentAt: time.Now().Format("02-01-2006 15:04:05")}
	}
}

func broadcast() {
	for newMsg := range broadcastCh {
		for client := range clients {
			msg, err := json.Marshal(newMsg)

			if err != nil {
				log.Println(err)
				return
			}
			client.conn.Write(client.ctx, websocket.MessageText, msg)
		}
	}
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("./view")))

	http.HandleFunc("/ws", wsHandler)

	http.HandleFunc("/clients", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		var res []*Client
		for c := range clients {
			res = append(res, c)
		}
		json.NewEncoder(w).Encode(res)
	})

	fmt.Println("Server on  :0")
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}

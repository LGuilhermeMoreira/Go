package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

type Client struct {
	nickname string
	ws       *websocket.Conn
	loged    bool
}

type Server struct {
	clients []Client
}

func (s *Server) handleSimpleWS(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("nickname")
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		// pular a verificação de segurança
		InsecureSkipVerify: true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	client := Client{
		nickname: value,
		ws:       conn,
		loged:    true,
	}

	s.clients = append(s.clients, client)

	// informando para todos os clientes que ele entrou
	for _, i := range s.clients {
		err = i.ws.Write(r.Context(), websocket.MessageText, []byte(i.nickname+" entrou"))
		if err != nil {
			log.Fatal(err)
			break
		}
	}

	// mandando mensagem para todos os clientes conectados

	for {
		_, data, err := conn.Read(r.Context())

		if err != nil {
			log.Fatalln(err)
			break
		}
		log.Printf("%v %v", r.RemoteAddr, string(data))

		// send messages to all contacts
		for _, i := range s.clients {
			msg := []byte("Mensagem generica")
			err = i.ws.Write(r.Context(), websocket.MessageText, msg)
			if err != nil {
				log.Fatal(err)
				break
			}
		}
	}
}

func (s *Server) handleLen(w http.ResponseWriter, r *http.Request) {
	tamanho := len(s.clients)

	msg := map[string]interface{}{
		"tamanho": tamanho,
	}

	resp, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, "Erro marshaling reponse", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Fatal(w.Write(resp))
}

func NewServer() *Server {
	return &Server{}
}

func main() {
	server := NewServer()

	http.HandleFunc("/", server.handleSimpleWS)

	http.HandleFunc("/len", server.handleLen)
	fmt.Println("Server on :0")
	log.Fatal(http.ListenAndServe(":3330", nil))
}

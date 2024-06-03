package main

import (
	"encoding/json"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

type Server struct {
	Clients map[*websocket.Conn]bool
}

func (s *Server) handleSimpleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		// pular a verificação de segurança
		InsecureSkipVerify: true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	s.Clients[conn] = true

	for client := range s.Clients {
		msg := []byte("Novo cliente cadastrado")
		client.Write(r.Context(), websocket.MessageText, msg)
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
		for client := range s.Clients {

			msg := []byte("Mensagem genérica")

			err = client.Write(r.Context(), websocket.MessageText, msg)

			if err != nil {
				log.Fatalln(err)
				break
			}
		}
	}
}

func (s *Server) handleLen(w http.ResponseWriter, r *http.Request) {
	tamanho := len(s.Clients)

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
	return &Server{
		Clients: make(map[*websocket.Conn]bool),
	}
}

func main() {
	server := NewServer()

	http.HandleFunc("/", server.handleSimpleWS)

	http.HandleFunc("/len", server.handleLen)

	log.Fatal(http.ListenAndServe(":3330", nil))
}

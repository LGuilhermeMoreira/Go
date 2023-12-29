package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	//mux := http.NewServeMux()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// pegando o context da requisição
	ctx := r.Context()
	// log imprime no comand line stdout
	log.Println("Request iniciada")
	defer log.Println("request finalizada")

	select {
	case <-ctx.Done():
		// log imprime no comand line stdout
		log.Println("request processada com fracasso")
		// imprime no browser
		w.Write([]byte("request processada com fracasso"))
	case <-time.After(5 * time.Second):
		// log imprime no comand line stdout
		log.Println("request processada com sucesso")
		// imprime no browser
		w.Write([]byte("request processada com sucesso"))
	}
}

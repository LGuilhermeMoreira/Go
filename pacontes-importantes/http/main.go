package main

import "net/http"

func main() {
	http.HandleFunc("/", BuscaCEP) // criando um end point

	// pode usar clousures
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Ola mundo"))
	// })
	http.ListenAndServe(":8000", nil) // abrindo || rodando um http servewr
}

func BuscaCEP(w http.ResponseWriter, r *http.Request) {
	// retorno da função associado ao "/"
	w.Write([]byte("Hello, World!"))
}

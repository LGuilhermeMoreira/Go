package main

import (
	"log"
	"net/http"
)

func main() {
	// Cria um servidor de arquivos para servir conteúdo estático da pasta "./public"
	fileServer := http.FileServer(http.Dir("./public"))

	// Cria um novo multiplexador (roteador) que será usado para registrar manipuladores de rota
	mux := http.NewServeMux()

	// Registra o servidor de arquivos para a rota raiz ("/")
	mux.Handle("/", fileServer)

	// Registra uma função de manipulação para a rota "/blog" que responde com "Ola blog"
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ola blog"))
	})

	// Inicia o servidor na porta 8000 e loga qualquer erro que ocorra durante a inicialização
	log.Fatal(http.ListenAndServe(":8000", mux))
}

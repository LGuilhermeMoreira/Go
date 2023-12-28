package main

import "net/http"

func main() {
	// Cria um novo multiplexador (roteador)
	mux := http.NewServeMux()

	// Registra a função HomeHandler para a rota raiz ("/")
	mux.HandleFunc("/", HomeHandler)

	// Registra a struct blog como um handler para a rota "/blog"
	mux.Handle("/blog", &blog{title: "My Blog"})

	// Inicia o servidor na porta 1000 usando o multiplexador criado
	http.ListenAndServe(":1000", mux)
}

// HomeHandler é um handler para a rota raiz ("/"). Responde com "ola mundo".
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ola mundo"))
}

// Struct blog com um campo "title".
type blog struct {
	title string
}

// ServeHTTP é implementado para satisfazer a interface http.Handler.
// Responde com o título da struct blog.
func (b *blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}

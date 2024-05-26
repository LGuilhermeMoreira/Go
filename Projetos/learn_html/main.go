package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	// Template para o formulário
	tmpl, err := template.ParseFiles("./html/form.html")
	if err != nil {
		panic(err)
	}

	// Manipulador para exibir o formulário
	mux.HandleFunc("GET /view/form", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := os.ReadFile("./html/form.html")

		if err != nil {
			http.Error(w, "Erro ao gerar o formulário", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, bytes) // Pass data to template if needed

		if err != nil {
			http.Error(w, "Erro ao gerar o formulário", http.StatusInternalServerError)
			return
		}
	})

	// Manipulador para capturar dados do formulário
	mux.HandleFunc("POST /view/form", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erro ao analisar o formulário", http.StatusBadRequest)
			return
		}

		nome := r.FormValue("nome")
		email := r.FormValue("email")

		// Validação e processamento de dados (opcional)
		// ...

		// Exibir dados capturados no terminal
		fmt.Println("Dados recebidos:")
		fmt.Println("Nome:", nome)
		fmt.Println("Email:", email)

		// (Opcional) Exibir dados capturados no template
		// err = tmpl.Execute(w, map[string]string{
		//     "nome": nome,
		//     "email": email,
		//     // "mensagem": mensagem de confirmação ou informação relevante
		// })
		// if err != nil {
		//     http.Error(w, "Erro ao gerar o formulário", http.StatusInternalServerError)
		//     return
		// }
	})

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", mux)
}

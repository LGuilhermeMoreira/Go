package main

import (
	"os"
	"text/template"
)

// Curso é uma struct que representa um curso com Nome e CargaHoraria.
type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	templates := []string{
		"header.html",
		".html",
		"footer.html",
	}
	// Define um template de texto usando a package "text/template"
	t := template.Must(template.New("index.html").ParseFiles(templates...))

	// Executa o template, passando os dados do curso e escreve a saída para os.Stdout
	err := t.Execute(os.Stdout, Cursos{
		{"c", 100},
		{"c++", 150},
		{"c#", 200},
	})

	// Verifica se houve algum erro durante a execução do template
	if err != nil {
		panic(err)
	}
}

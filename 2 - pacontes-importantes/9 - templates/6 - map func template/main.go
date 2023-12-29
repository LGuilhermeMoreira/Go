package main

import (
	"os"
	"strings"
	"text/template"
)

// Curso é uma struct que representa um curso com Nome e CargaHoraria.
type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {

	templates := []string{
		"header.html",
		"index.html",
		"footer.html",
	}

	t := template.New("index.html")
	//t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper}) utilizando a função da lib strings
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))

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

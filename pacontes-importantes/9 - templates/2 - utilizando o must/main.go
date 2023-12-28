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

func main() {
	// Cria uma instância da struct Curso
	curso := Curso{Nome: "Go", CargaHoraria: 48}

	// Define um template de texto usando a package "text/template"
	t := template.Must(template.New("cursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))

	// Executa o template, passando os dados do curso e escreve a saída para os.Stdout
	err := t.Execute(os.Stdout, curso)

	// Verifica se houve algum erro durante a execução do template
	if err != nil {
		panic(err)
	}
}

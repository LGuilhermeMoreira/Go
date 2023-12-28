package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{Nome: "Go", CargaHoraria: 48}

	// criando um template
	tmp := template.New("CursoTemplate")

	// Analizando o padrão do template
	tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")

	// esse os.Stdout é utilizado para imprimir no terminal, mas suporta qualquer coisa que tenha um metodo Write passando um []byte() como parametro
	err := tmp.Execute(os.Stdout, curso)

	if err != nil {
		panic(err)
	}
}

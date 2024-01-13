package main

import (
	"net/http"
	"text/template"
)

// Curso é uma struct que representa um curso com Nome e CargaHoraria.
type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("cursos").Parse("Curso: {{.Nome}} Carga Horária: {{.CargaHoraria}} horas"))

		// como w tem uma função Write então cabe como primeiro parametro
		// no caso o retorno é um w.Write([]bytes()), onde esse []bytes() == o template passado
		err := t.Execute(w, Curso{"FullCycle Go", 40})

		if err != nil {
			panic(err)
		}

	})

	println("Server Online!")
	// subindo o http Server
	http.ListenAndServe(":8000", nil)
	println("Server Offline!")
}

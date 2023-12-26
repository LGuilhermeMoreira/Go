package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("https://pokeapi.co/api/v2/pokemon/ditto")
	if err != nil {
		panic(err)
	}
	// defer atrasa a execução, executa sempre no fim do programa
	defer request.Body.Close()

	response, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(response))

	//request.Body.Close()
}

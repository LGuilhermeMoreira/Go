package main

import (
	"fmt"
	"net/http"
	"pokeapi/structs"
)

func sendErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// digite o nome dos pokemons
	print("Digite o nome do Pokemon que deseja ")
	var input string

	pokemon := structs.Pokemon{}

	print(pokemon)

	_, err := fmt.Scan(&input)
	// verificando se tem algum erro de input
	sendErr(err)

	fmt.Printf("meu nome é: %s\n", input)

	request, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + input)
	sendErr(err)

	defer request.Body.Close()

}

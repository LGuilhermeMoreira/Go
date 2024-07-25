package main

import (
	"fmt"
	"os"
	"test_lenguage/pkg"
)

func main() {

	bytes, err := os.ReadFile("input.Alang")

	if err != nil {
		panic(err)
	}

	input := string(bytes)

	tokens, err := pkg.Lex(input)

	if err != nil {
		fmt.Println("ERRO")
	} else {
		for _, token := range tokens {
			fmt.Printf("%s (%s)\n", token.Type, token.Lexeme)
		}
	}
}

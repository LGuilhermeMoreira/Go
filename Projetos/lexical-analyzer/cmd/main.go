package main

import (
	"flag"
	"fmt"
	"os"
	"test_lenguage/pkg"
)

func main() {

	str := flag.String("path", "", "path to file .Alang")
	flag.Parse()

	if *str == "" {
		panic("flag path is empty")
	}

	bytes, err := os.ReadFile(*str)

	if err != nil {
		panic(err)
	}

	input := string(bytes)

	tokens, err := pkg.Lex(input)

	if err != nil {
		fmt.Printf("ERRO: %v\n", err)
	} else {
		for _, token := range tokens {
			fmt.Printf("%s (%s)\n", token.Type, token.Lexeme)
		}
	}
}

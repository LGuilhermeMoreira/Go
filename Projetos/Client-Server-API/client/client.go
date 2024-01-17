package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("localhost:8080/conversao")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	dados, err := io.ReadAll(response.Body)

	fmt.Println(string(dados))
}

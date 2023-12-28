package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] {
		println(url)

		request, err := http.Get(url)

		if err != nil {
			fmt.Println(os.Stderr, "erro ao fazer requisição: %v\n", err)
		}

		defer request.Body.Close()

		response, err := io.ReadAll(request.Body)

		if err != nil {
			fmt.Println(os.Stderr, "erro ao ler request: %v\n", err)
		}

		var data ViaCEP

		err = json.Unmarshal(response, &data)

		if err != nil {
			fmt.Println(os.Stderr, "Erro ao fazer parse: %v\n", err)
		}

		fmt.Println(data)

	}
}

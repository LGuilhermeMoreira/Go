package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	// Cria um cliente HTTP com as configurações padrão.
	c := http.Client{}

	// Cria um buffer contendo dados JSON para serem enviados na requisição.
	jsonVar := bytes.NewBuffer([]byte(`{"name" : "weslay"}`))

	// Faz uma requisição POST para o URL especificado (http://google.com) usando o cliente HTTP criado.
	// O conteúdo da requisição é especificado como JSON no cabeçalho "Content-Type".
	resp, err := c.Post("http://google.com", "application/json", jsonVar)

	// Verifica se ocorreu algum erro durante a requisição.
	if err != nil {
		// Em caso de erro, interrompe a execução do programa e exibe o erro.
		panic(err)
	}

	// Garante que o corpo da resposta (response body) será fechado no final da função main.
	defer resp.Body.Close()

	// Copia o corpo da resposta para a saída padrão (no caso, os.Stdout).
	// Pode ser usado para imprimir o conteúdo da resposta no terminal.
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}

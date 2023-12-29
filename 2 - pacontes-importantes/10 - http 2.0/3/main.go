package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	// Cria um cliente HTTP com as configurações padrão.
	c := http.Client{}

	// Cria um novo objeto de solicitação (request) para a URL especificada (http://google.com).
	req, err := http.NewRequest("GET", "http://google.com", nil)

	// Verifica se ocorreu algum erro ao criar a solicitação.
	if err != nil {
		// Em caso de erro, interrompe a execução do programa e exibe o erro.
		panic(err)
	}

	// Define o cabeçalho da solicitação para indicar que o cliente aceita conteúdo JSON.
	req.Header.Set("Accept", "application/json")

	// Faz a solicitação usando o cliente HTTP criado e a solicitação preparada.
	resp, err := c.Do(req) // junta o clinete http com o objeto de request

	// Verifica se ocorreu algum erro durante a solicitação.
	if err != nil {
		// Em caso de erro, interrompe a execução do programa e exibe o erro.
		panic(err)
	}

	// Garante que o corpo da resposta (response body) será fechado no final da função main.
	defer resp.Body.Close()

	// Lê o corpo da resposta e armazena em 'body'.
	body, err := ioutil.ReadAll(resp.Body)

	// Verifica se ocorreu algum erro durante a leitura do corpo da resposta.
	if err != nil {
		// Em caso de erro, interrompe a execução do programa e exibe o erro.
		panic(err)
	}

	// Converte o corpo da resposta para uma string e a imprime.
	println(string(body))
}

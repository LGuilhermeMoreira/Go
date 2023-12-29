package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// Cria um cliente HTTP com um tempo limite de 1 segundo para a requisição.
	// Esse Timeout representa o tempo máximo que a requisição pode durar antes de ser cancelada.
	c := http.Client{Timeout: time.Second}

	// Alternativamente, você pode usar um tempo limite mais curto, por exemplo, em milissegundos.
	// c := http.Client{Timeout: time.Millisecond}

	// Faz uma requisição GET para o URL especificado (http://google.com) usando o cliente HTTP criado.
	resp, err := c.Get("http://google.com")

	// Verifica se ocorreu algum erro durante a requisição.
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

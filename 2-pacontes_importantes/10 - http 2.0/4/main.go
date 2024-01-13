package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

/*
PACOTE DE CONTEXTO
PERMITE QUE PASSEMOS AS INFORMAÇÕES DELES PARA DIVERSAS CHAMADAS NO NOSSO SISTEMA E
, PODEMOS CANCELAR ESSE CONTEXTO. QUANDO CANCELADO, ELE É PARADO NA HORA
*/

func main() {
	// por padrão ctx é usado para contexto
	//context.Barckground() é usado para criar um novo contexto (contexto vazio)
	ctx := context.Background()
	// criando um contexto com timeout -> caso passe de 5 seg ele irar parar
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second) // o 5 representa os 5 segundos
	defer cancel()                                         // a função cancel CANCELA o contexto

	//poderia criar um contexto que so cacelaria com a chamada da função cancel
	//ctx,cancel := context.WithCancel(ctx)

	// criando uma requisição http utilizando contexto
	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	// utilizando um client
	/*
		c : = http.Cliente()

		resp,err := c.Do(req)
	*/

	// executando sem um client
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}

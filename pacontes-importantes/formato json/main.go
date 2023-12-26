package main

import (
	"encoding/json"
	"fmt"
)

type Conta struct {
	Numero int
	Saldo  int
}

func main() {
	conta := Conta{Numero: 198019, Saldo: 2000}

	// função Marshal transforma em JSON, e tem o retorno me bytes
	request, err := json.Marshal(conta)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(request))

}

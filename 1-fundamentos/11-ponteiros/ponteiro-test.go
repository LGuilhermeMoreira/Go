package main

import "fmt"

func somaNormal(a, b int) int {
	return a + b
}

func somaDiferenteComPonteiros(a, b *int) int {
	*a = *a + *b
	*b = *b + *a

	return *a + *b
}

type Pessoa struct {
	nome string
}

func (p *Pessoa) nomeTrocado() {
	p.nome = "Nome trocado"
}

func main() {
	jp := Pessoa{nome: "Jp"}
	jp.nomeTrocado()

	a := 5
	b := 9

	fmt.Printf("valores iniciais de a e b: %d %d", a, b)

	fmt.Printf("soma normal de a e b: %d", somaNormal(a, b))

	fmt.Printf("soma diferente com ponteiros a e b: %d", somaDiferenteComPonteiros(&a, &b))

}

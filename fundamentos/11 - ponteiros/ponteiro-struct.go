package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}

func (p Pessoa) andou() {
	p.nome += " andou"
	fmt.Printf("A pessoa %s andou", p.nome)
}

func (p *Pessoa) andou_ponteiro() {
	p.nome += " andou"
	fmt.Printf("A pessoa %s andou\n", p.nome)
}

func newPessoa() *Pessoa {
	return &Pessoa{nome: "Nova pessoa", idade: 0}
}

func (p *Pessoa) caminhou() {
	p.nome += " caminhou"
}

func main() {
	// jp := Pessoa{nome: "jp",
	// 	idade: 12}

	// vitu := Pessoa{nome: "vitor", idade: 15}
	// jp.andou()

	// vitu.andou_ponteiro()

	// fmt.Printf("novo nome de vitu: %s\n", vitu.nome)

	a := newPessoa()
	b := newPessoa()
	a.nome = "jose"
	b.caminhou()
	fmt.Printf("nome de a: %s\nnome de b: %s\n", a.nome, b.nome)

}

package main

import "fmt"

// Definição da struct Endereco, representando um endereço com rua e número.
type Endereco struct {
	rua    string
	numero int
}

// Definição da struct Cliente, representando um cliente com nome, idade, estado de ativo e endereço.
type Cliente struct {
	Nome   string   // Nome do cliente.
	Idade  int      // Idade do cliente.
	Ativo  bool     // Estado de ativo do cliente.
	Endero Endereco // Incorporação da struct Endereco para representar o endereço do cliente.
}

// Método Desativar, que desativa o cliente (define Ativo como false) e imprime uma mensagem.
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Println("Cliente desativado")
}

// Método show, que imprime as informações do cliente.
func (c Cliente) show() {
	fmt.Printf("Nome: %s\n", c.Nome)
	fmt.Printf("Idade: %d\n", c.Idade)
	fmt.Printf("Ativo: %t\n", c.Ativo)
	fmt.Printf("Rua: %s\n", c.Endero.rua)
	fmt.Printf("Numero: %d\n", c.Endero.numero)
}

// Definição de uma interface Pessoa, que exige a implementação do método Desativar.
type Pessoa interface {
	Desativar()
}

// Função DesativarPessoa, que aceita um objeto que implementa a interface Pessoa e chama o método Desativar.
func DesativarPessoa(p Pessoa) {
	p.Desativar()
}

func main() {
	jorge := Cliente{
		Nome:  "Jorge",
		Idade: 30,
		Ativo: true,
		Endero: Endereco{
			rua:    "Pedro Cabral de Oliveria",
			numero: 2054,
		},
	}

	// posso chamar o metodo desativarpessoa passando um cliente,
	// pois o cliente implementa a interface pessoa

	DesativarPessoa(jorge)
}

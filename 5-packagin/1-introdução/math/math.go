package math

import "fmt"

// tudo que começa maiusculo é Publico

// tudo que começa minusculo é privado

// TUDO PUBLICO
type Math struct {
	A int
	B int
}

func (m *Math) Add() int {
	return m.A + m.B
}

// TUDO PRIVADO
type mathPrivada struct {
	a int
	b int
}

func (m mathPrivada) add() int {
	return m.a + m.b
}

func (m mathPrivada) ExecuteAdd() {
	fmt.Println(m.add())
}

// para ter acesso a essa struct é nessário criar uma função PUBLICA que retorne ela
func NewmathPrivada(a, b int) mathPrivada {
	return mathPrivada{a: a, b: b}
}

// PARA TRABALHAR COM ATRIBUTOS PRIVADOS
type User struct {
	name string
	id   int
}

func NewUser(name string, id int) User {
	return User{name: name, id: id}
}

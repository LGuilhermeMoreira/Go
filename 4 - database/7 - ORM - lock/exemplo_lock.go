package main

import (
	"fmt"
	"sync"
	"time"
)

// Dados representa a estrutura de dados compartilhada
type Dados struct {
	valor  int
	mu     sync.Mutex // Mutex para lock pessimista
	versao int        // Versão para lock otimista
}

// LockPessimista realiza um lock pessimista para modificar os dados
func LockPessimista(d *Dados, novoValor int) {
	d.mu.Lock()
	defer d.mu.Unlock()

	// Operações nos dados
	d.valor = novoValor
	fmt.Println("Dados modificados pessimisticamente.")
}

// LockOtimista realiza um lock otimista para modificar os dados
func LockOtimista(d *Dados, novoValor int) {
	// Tentativa de obter um lock otimista
	versaoAntiga := d.versao
	d.versao++

	// Simula uma verificação de conflito
	time.Sleep(time.Millisecond * 100)

	// Verifica se ocorreu algum conflito durante a operação
	if versaoAntiga != d.versao-1 {
		fmt.Println("Conflito detectado. Operação cancelada.")
		return
	}

	// Operações nos dados
	d.valor = novoValor
	fmt.Println("Dados modificados optimisticamente.")
}

func test_look() {
	dados := Dados{valor: 42}

	// Exemplo de lock pessimista
	go LockPessimista(&dados, 99)
	go LockPessimista(&dados, 88)

	time.Sleep(time.Second) // Aguarda a conclusão das goroutines

	// Exemplo de lock otimista
	go LockOtimista(&dados, 77)
	go LockOtimista(&dados, 66)

	time.Sleep(time.Second) // Aguarda a conclusão das goroutines
}

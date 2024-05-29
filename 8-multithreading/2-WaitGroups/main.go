package main

import (
	"fmt"
	"sync"
	"time"
)

func task(key string, wg *sync.WaitGroup) {
	for i := range 10 {
		fmt.Printf("%d - task: %v\n", i, key)
		time.Sleep(1 * time.Second)
		wg.Done() // consome o numero de opreções que foram criadas
	}
}

// thread 1
func main() {
	// determina
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25) // numero de operações que vão rodar no total
	// thread 2
	go task("A", &waitGroup)
	// thread 3
	go task("B", &waitGroup)

	// thread 4
	// declaração de uma go routine anonima
	go func(key string) {
		for i := 0; i < 10; i++ {
			fmt.Println("Função anonima rodando!", key)
			waitGroup.Done()
		}
	}("Key")

	// enquanto essas operações não foram concluidas, ele segura o programa
	waitGroup.Wait()
	fmt.Printf("Fim do processamento")
}

// video 8

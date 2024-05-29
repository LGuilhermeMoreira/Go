package main

import (
	"fmt"
	"time"
)

func task(key string) {
	for i := range 10 {
		fmt.Printf("%d - task: %v\n", i, key)
		time.Sleep(1 * time.Second)
	}
}

// thread 1
func main() {
	// thread 2
	go task("A")
	// thread 3
	go task("B")

	// thread 4
	// declaração de uma go routine anonima
	go func(key string) {
		for i := 0; i < 10; i++ {
			fmt.Println("Função anonima rodando!", key)
		}
	}("Key")

	// necessita de algo mantendo o programa vivo para as go routines rodarem
	time.Sleep(time.Second * 15)
	fmt.Printf("Fim do processamento")
}

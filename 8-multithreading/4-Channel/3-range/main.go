package main

import "fmt"

func main() {
	channel := make(chan int)
	go publish(channel)
	// essa função não roda em background para n matar a execução do thread acima
	//
	read(channel)
}

func read(ch chan int) {
	for i := range ch {
		fmt.Printf("recived: %d\n", i)
	}
}

func publish(ch chan int) {
	for i := range 10 {
		ch <- i
	}
	// dando um close no channel, evita o caso de deadlock
	close(ch)
}

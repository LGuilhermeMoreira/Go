package main

import "fmt"

/*
	Um channel é responsável por comunicação entre threads

	A passagem de informação só é feita quando o canal está vazio

*/

// thread 1
func main() {

	channel := make(chan string)

	// thread 2
	go func() {
		channel <- "Olá mundo"
	}()

	msg := <-channel

	fmt.Println(msg)
}

package main

import "fmt"

func main() {
	hello := make(chan string)
	go recebe("Alou", hello)
	passar(hello)
}

// receive only
// dessa meneira você restringe o canal para apenas receber informção
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// send Only
// dessa maneira você restringe o canal para apenas passar informação
func passar(data <-chan string) {
	fmt.Println(<-data)
}

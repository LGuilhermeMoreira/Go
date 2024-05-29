package main

func main() {
	// é como se o canal agora podesse guardar 2 objetos
	ch := make(chan string, 2)
	//objeto 1
	ch <- "Hello"
	//objeto 2
	ch <- "World"

	println(<-ch)
	println(<-ch)
}

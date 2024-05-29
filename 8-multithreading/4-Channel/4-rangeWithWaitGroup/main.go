package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)
	go publish(channel)

	go read(channel, &wg)
	wg.Wait()
}

func read(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Printf("recived: %d\n", i)
		wg.Done()
	}
}

func publish(ch chan int) {
	for i := range 10 {
		ch <- i
	}
	close(ch)
}

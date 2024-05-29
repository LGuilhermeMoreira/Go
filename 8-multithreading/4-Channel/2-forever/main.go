package main

import "fmt"

func main() {
	/*
		Essa abordagem de criar um canal e esvaziar ele,
		segura as routines, entretando causa deadlock
	*/
	forever := make(chan bool)

	go func() {

		for i := range 10 {
			fmt.Println(i)
		}

		// para evitar o deadlock preenchemos o canal
		forever <- true
	}()

	<-forever
}

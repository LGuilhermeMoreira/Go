package main

import "fmt"

type ID int

var (
	a bool    = false
	b int     = 12
	c float64 = 23.5
	d string  = "uma string"
	e ID      = 12
)

func main() {
	var meuArray [3]int
	meuArray[0] = 1

	meuArray[1] = 2

	meuArray[2] = 3

	for i, v := range meuArray {
		fmt.Printf("indice: %v valor: %v \n", i, v)
	}

}

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
	// trazendo o tipo
	fmt.Printf("==============================\n")
	fmt.Printf("O tipo é %T\n", a)
	fmt.Printf("O tipo é %T\n", b)
	fmt.Printf("O tipo é %T\n", c)
	fmt.Printf("O tipo é %T\n", d)
	fmt.Printf("O tipo é %T\n", e)

	fmt.Printf("==============================\n")
	// trazendo o valor
	fmt.Printf("O valor é %v\n", a)
	fmt.Printf("O valor é %v\n", b)
	fmt.Printf("O valor é %v\n", c)
	fmt.Printf("O valor é %v\n", d)
	fmt.Printf("O valor é %v\n", e)
	fmt.Printf("==============================\n")
}

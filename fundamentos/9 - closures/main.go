package main

import (
	"fmt"
)

// caso queria utilizar um  número de parametros que eu n saiba calcular
func somaInfinita(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func main() {
	sumResult_3 := somaInfinita(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Printf("%d\n", sumResult_3)

	total := func(i int) int {
		return somaInfinita(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) * i
	}

	fmt.Printf("%d\n", total(8*2))
}

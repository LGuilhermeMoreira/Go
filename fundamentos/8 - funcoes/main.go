package main

import (
	"errors"
	"fmt"
)

// func <nome da função>(<parametros>)<tipos de retorno> {<código>}
func sum_1(a int, b int) int {
	return a + b
}

func sumWithBool(a int, b int) (int, bool) {
	if a%2 == 0 && b%2 == 0 {
		return a + b, true
	}
	return a + b, false
}

// não existe try e catch em go então podemos adicionar mais um retorno como erro
func sum_2(a int, b int) (int, error) {
	if a+b > 50 {
		return 0, errors.New("error: soma maior que 50 não pode")
	}
	return a + b, nil
}

// caso queria utilizar um  número de parametros que eu n saiba calcular
func somaInfinita(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func main() {
	fmt.Printf("%d\n", sum_1(4, 2))
	sumResult, isEven := sumWithBool(4, 2)
	fmt.Printf("%d %t\n", sumResult, isEven)
	sumResult_2, errors := sum_2(50, 1)
	fmt.Printf("%d %v\n", sumResult_2, errors)
	sumResult_3 := somaInfinita(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Printf("%d\n", sumResult_3)
}

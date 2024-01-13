package main

import "fmt"

func main() {
	x := 10
	y := 5

	if x > 10 && y > 5 {
		fmt.Println("Ambas as condições são verdadeiras.")
	} else if x == 10 || y == 5 {
		fmt.Println("Pelo menos uma das condições é verdadeira.")
	} else {
		fmt.Println("Nenhuma das condições é verdadeira.")
	}

	if x > y {
		fmt.Println("x é maior que y.")
	} else if x < y {
		fmt.Println("x é menor que y.")
	} else {
		fmt.Println("x é igual a y.")
	}
}

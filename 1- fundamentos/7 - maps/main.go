package main

import "fmt"

func main() {
	salarios := map[string]int{"Weslay": 1000, "João": 2000, "Ana": 3000}

	fmt.Printf("salarios: %v\n", salarios)

	delete(salarios, "Weslay")
	fmt.Printf("salarios: %v\n", salarios)

	salarios["Wagner"] = 4000
	fmt.Printf("salarios: %v\n", salarios)

	// criando um map
	sal_1 := make(map[string]int)
	sal_2 := map[string]int{}

	fmt.Printf("salarios: %v\nsalarios: %v\n", sal_1, sal_2)

	// a função range retorna o indice e o valor

	for nome, salarios := range salarios {
		fmt.Printf("%v : %d\n", nome, salarios)
	}

	//para ignorar um valor utilizamos o _
	/*			EX -  ignorando o nome
	for _,salarios := range salarios {
		fmt.Printf("%v : %d\n", nome, salarios)
	}

	*/

	str := "iai comparsa"

	for _, caracter := range str {
		fmt.Printf("%c ", caracter)
	}

}

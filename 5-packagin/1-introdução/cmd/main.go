package main

import (
	"fmt"
	"packagin/1/math"
	"packagin/1/test1"
)

func main() {

	m := math.Math{
		A: 12,
		B: 233,
	}

	fmt.Println(m.Add())

	test1.ABCDTest()

	fmt.Println(test1.Variavel)

	// consigo criar um User vazio, mas não consigo acessar os atributos
	u := math.User{}
	fmt.Println(u)

	// para criar com atributos, é necessário utilizar a função NewUser

	u = math.NewUser("ronaldo", 123)

	fmt.Println(u)

	mP := math.NewmathPrivada(1, 2)

	mP.ExecuteAdd()
}

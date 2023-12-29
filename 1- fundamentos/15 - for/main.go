package main

func main() {
	// go tem apenas o for

	//for comum
	for i := 0; i < 10; i++ {
		println(i)
	}

	//for com range

	numeros := []int{1, 3, 4, 5, 6}

	// função range retorna o indice e o valor
	for indice, valor := range numeros {
		println(indice, valor)
	}

	// utilizando o for como um while
	k := 0
	for k < 10 {
		println(k)
		k++
	}

	// loop infinito
	for {
		println("abcd")
	}
}

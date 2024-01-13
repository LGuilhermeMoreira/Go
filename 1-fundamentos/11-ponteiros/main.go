package main

// memoria -> endereço -> valor
var a int = 10

var ponteiro *int = &a // ponteiro para inteiro

var b *int = &a

func show() {
	println(a, ponteiro, *ponteiro, b, *b)
}

func soma(a, b int) int {
	return a + b
}

func soma_com_ponteiro(a, b *int) int {
	*a = 2
	*b = 3
	return *a + *b
}

func most(b *int) {
	*b = 12
}

func main() {
	//show()
	a_1 := 166
	// a_2 := 20
	// b_1 := 4
	// b_2 := 5
	//print(soma(b_1, b_2))
	//print(soma_com_ponteiro(&a_1, &a_2))
	println(a_1)
	most(&a_1)
	println(a_1)
}

package main

func main() {

	var minhaVar interface{} = "weslay"
	println(minhaVar.(string)) // pegando o verdadeiro valor da string

	// tentando converter o valor da interface vazia para int
	valor, ok := minhaVar.(int)

	println(valor, ok)
}

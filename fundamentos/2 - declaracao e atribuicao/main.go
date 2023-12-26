package main

// escopo global
// constantes

const a = "Olá, Mundo!"

var (
	b int     // int
	c bool    // boolean
	d string  // string
	e float64 // float <mt estranho esse float64>
	f string  = "letra f"
)

func main() {

	// criando e declarando variavel
	g := 12

	g = 234

	h := false

	h = false

	println(b)
	println(c)
	println(d)
	// chamando função
	x()

	println(g)

	println(h)
}

func x() {
	print("abcd_efgh\n")
}

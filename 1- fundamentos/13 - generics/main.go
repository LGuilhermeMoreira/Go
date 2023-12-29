package main

type MyNumber int

func Soma_int(m map[string]int) int {
	soma := 0
	for _, v := range m {
		soma += v
	}
	return soma
}

func Soma_float64(m map[string]float64) float64 {
	soma := 0.0
	for _, v := range m {
		soma += v
	}
	return soma
}

// criando uma função que aceita tanto um float como int
func Soma_generics[T int | float64](m map[string]T) T {
	var soma T = 0
	for _, v := range m {
		soma += v
	}
	return soma
}

func Soma_generics_constraint[T Number](m map[string]T) T {
	var soma T = 0
	for _, v := range m {
		soma += v
	}
	return soma
}

// criando uma constraint que serve como int ou
type Number interface {
	//~int | float64 // adiciona um ~ para que qualquer valor que seja um inteiro (MyNumber) seja aceito

	// para evitar erros, sempre adiciona ~ me tudo
	~int | ~float64
}

func Compara[T Number](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"guilherme": 12, "jaime": 15}
	n := map[string]float64{"guilherme": 12.2, "jaime": 15.1}

	print("================================")
	print(Soma_int(m))
	print("================================")
	print(Soma_float64(n))
	print("================================")
	print("================================")
	print(Soma_generics[int](m))
	print(Soma_generics[float64](n))
	print("================================")
	print("================================")

	o := map[string]MyNumber{"guilherme": 12, "jaime": 15}

	print(Soma_generics_constraint(o))

}

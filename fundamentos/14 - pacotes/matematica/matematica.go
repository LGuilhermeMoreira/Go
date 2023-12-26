package matematica

func Soma[T int | float64](x, y T) T {
	return x + y
}

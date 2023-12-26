package main

import "fmt"

// uma interface vazia permite que qualquer tipo seja aceito

func main() {
	var x interface{} = 10
	var y interface{} = "iai comparca"
	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("tipo: %T valor: %v\n", t, t)
}

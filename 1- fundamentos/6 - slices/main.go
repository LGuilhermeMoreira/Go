package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4}

	// entendendo slices
	fmt.Println("===========================================================")
	fmt.Printf("tamanho: %d\ncapacidade: %d\narray: %v\n", len(array), cap(array), array)
	fmt.Println("===========================================================")
	fmt.Printf("tamanho: %d\ncapacidade: %d\narray: %v\n", len(array[:0]), cap(array[:0]), array[:0])
	fmt.Println("===========================================================")
	fmt.Printf("tamanho: %d\ncapacidade: %d\narray: %v\n", len(array[:3]), cap(array[:3]), array[:3])
	fmt.Println("===========================================================")

	// adicionando o valor 5 no array
	array = append(array, 5)

	fmt.Println("valor adicionado no array")

	fmt.Println("===========================================================")
	fmt.Printf("tamanho: %d\ncapacidade: %d\narray: %v\n", len(array), cap(array), array)
	fmt.Println("===========================================================")
	fmt.Printf("tamanho: %d\ncapacidade: %d\narray: %v\n", len(array[:0]), cap(array[:0]), array[:0])
	fmt.Println("===========================================================")
	fmt.Printf("tamanho: %d\ncapacidade: %d\narray: %v\n", len(array[:3]), cap(array[:3]), array[:3])
	fmt.Println("===========================================================")
}

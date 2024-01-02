package main

import "github.com/google/uuid"

type Product struct {
	ID    string
	Name  string
	Price float64
}

// com essa função não teria o incremento do ID
// func NewProduct(id, name string, price float64) *Product {
// 	return &Product{
// 		ID:    id,
// 		Name:  name,
// 		Price: price,
// 	}
// }

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {

}

package main

/*
	COM O GO WORK
	é possível resolver o problema de ter 2 go.mod, entretando você não consegue
	utilizar o go mod tidy para acessar libs de outras pessoas. Ex: uuid
	SOLUÇÃO: go mod tidy -e
*/

import (
	"fmt"
	"packagin/4/math"

	"github.com/google/uuid"
)

func main() {
	m := math.NewUser("rodrigo", 2)
	fmt.Println(m)
	fmt.Println(uuid.New().String())
}

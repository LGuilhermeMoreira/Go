package main

// esse pacote está indisponivel pois ele não está presente
// no github ou em qualquer outro lugar para ser acessado
import (
	"fmt"
	"packagin/3/math"
)

/*
ENTRETANDO
com esse comando: go mod edit -replace packagin/3/math=../math e em seguida um go mod tidy
*/

func main() {
	m := math.NewUser("rodrigo", 2)
	fmt.Println(m)
}

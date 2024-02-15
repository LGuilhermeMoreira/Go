package main

import "estrutura_de_dados/tree"

func main() {
	// lista := list.Start()

	// for i := 0; i < 10; i++ {
	// 	lista.Add(i)
	// }

	// lista.Show()

	arvore := tree.Start()

	arvore.Add(12)
	arvore.Add(1)
	arvore.Add(54)

	arvore.Show()
}

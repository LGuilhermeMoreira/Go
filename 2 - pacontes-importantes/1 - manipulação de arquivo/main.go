package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// criando arquivo txt
	f, err := os.Create("arquivo.txt") // f é o arquivo e err são os erros

	// checando se não tem erro
	if err != nil {
		panic(err)
	}

	// caso o conteudo seja uma string
	//tamanho, err := f.WriteString("Ola mundo!") // tamanho é o tamanho do arquivo

	// caso o conteudo seja outra coisa
	tamanho, err := f.Write([]byte("escrevendo no arquivo")) // tamanho é o tamanho do arquivo

	if err != nil {
		panic(err)
	}

	fmt.Printf("arquvio criado com sucesso! tamanho: %d bytes\n", tamanho)
	f.Close()

	//leitura
	arquivo, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	//leitura de linha a linha
	arquivo2, err := os.Open("arquivo.txt")

	if err != nil {
		panic(err)
	}

	// importei a biblioteca bufio para ler linha por linha
	reader := bufio.NewReader(arquivo2)

	// delimitei quantos bytes ele vai ler por vez
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// remover arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}

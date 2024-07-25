# Analisador Léxico em Go

Este é um simples analisador léxico desenvolvido em Go. Ele lê um arquivo de código fonte com a extensão `.Alang` e analisa os tokens definidos.

## Pré-requisitos

Certifique-se de ter o Go instalado em sua máquina. Você pode baixar e instalar o Go a partir do site oficial: [golang.org](https://golang.org/)

## Estrutura do Projeto

├── cmd
│ └── main.go
├── input
│ └── a.Alang
├── pkg
│ └── lex.go
│ └── lex_test.go
├── go.mod
├── go.sum
└── README.md

## Exemplo de Comando para rodar o projeto

go run cmd/main.go -path input/a.Alang
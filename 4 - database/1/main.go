package main

import (
	// Usando a interface para o banco de dados.
	"database/sql"

	// Importando os drivers do banco de dados MySQL.
	_ "github.com/go-sql-driver/mysql"

	// Importando a biblioteca UUID do Google para gerar IDs únicos.
	"github.com/google/uuid"
)

func main() {

	//sql.Open("<banco de dados>","<user>:<password>@tcp(localhost:<porta>)/<nome do DB")
	// Para fazer a conexão com o banco de dados MySQL.
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")

	// Verifica se ocorreu algum erro ao abrir a conexão com o banco de dados.
	if err != nil {
		panic(err)
	}

	// Garante que a conexão com o banco de dados será fechada no final da função main.
	defer db.Close()

	// Cria uma instância de Product usando a função NewProduct, que gera um ID único.
	product := NewProduct("IPHONE 15", 9000.0)

	// Insere o produto no banco de dados.
	err = insertProduct(db, product)

	// Verifica se ocorreu algum erro durante a inserção do produto no banco de dados.
	if err != nil {
		panic(err)
	}
}

// Definindo a estrutura do produto.
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

// Função construtora para criar uma nova instância de Product com um ID único gerado.
func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(), // Gera um ID único usando a biblioteca UUID do Google.
		Name:  name,
		Price: price,
	}
}

// Função para inserir um produto no banco de dados.
func insertProduct(db *sql.DB, product *Product) error {
	// Prepara a declaração SQL para a inserção do produto.
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES (?,?,?)")

	// Verifica se ocorreu algum erro ao preparar a declaração SQL.
	if err != nil {
		return err
	}

	// Garante que a declaração SQL será fechada no final da função insertProduct.
	defer stmt.Close()

	// podemos pegar o result e checar algumas informações
	// result, err = stmt.Exec(product.ID, product.Name, product.Price)

	// Executa a declaração SQL para inserir o produto no banco de dados.
	_, err = stmt.Exec(product.ID, product.Name, product.Price)

	// Verifica se ocorreu algum erro durante a execução da declaração SQL.
	if err != nil {
		return err
	}

	// Retorna nil, indicando que a operação foi bem-sucedida.
	return nil
}

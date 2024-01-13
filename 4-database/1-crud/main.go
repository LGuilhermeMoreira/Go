package main

import (
	// Usando a interface para o banco de dados.
	"database/sql"
	"fmt"

	//standart input&output libary

	// Importando os drivers do banco de dados MySQL.
	_ "github.com/go-sql-driver/mysql"

	// Importando a biblioteca UUID do Google para gerar IDs únicos.
	"github.com/google/uuid"
)

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

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)

	if err != nil {
		return err
	}

	return nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func selectOneProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id,name,price from products where id = ?")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product

	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func showOneProduct(db *sql.DB, id string) error {
	product, err := selectOneProduct(db, id)

	if err != nil {
		return err
	}

	fmt.Printf("ID: %v\nName: %v\nPrice: %v\n", product.ID, product.Name, product.Price)
	return err
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	var products []Product

	stmt, err := db.Prepare("select id,name,price from products")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	for rows.Next() {
		var p Product

		err := rows.Scan(&p.ID, &p.Name, &p.Price)

		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func main() {

	//sql.Open("<banco de dados>","<user>:<password>@tcp(localhost:<porta>)/<nome do DB")
	// Para fazer a conexão com o banco de dados MySQL.
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")

	// // Verifica se ocorreu algum erro ao abrir a conexão com o banco de dados.
	if err != nil {
		panic(err)
	}

	// // Garante que a conexão com o banco de dados será fechada no final da função main.
	defer db.Close()

	// // Cria uma instância de Product usando a função NewProduct, que gera um ID único.
	//product := NewProduct("", 0.0)

	// // Insere o produto no banco de dados.
	//err = insertProduct(db, product)

	// Verifica se ocorreu algum erro durante a inserção do produto no banco de dados.
	// if err != nil {
	// 	panic(err)
	// }

	// exemplode id hashirizado = a7bfcfba-6e94-4326-a37f-bf7de1e399df
	var ID string = "a7bfcfba-6e94-4326-a37f-bf7de1e399df"

	//p, err := selectOneProduct(db, ID)

	if err != nil {
		panic(err)
	}

	err = showOneProduct(db, ID)

	if err != nil {
		panic(err)
	}

	productsList, err := selectAllProducts(db)

	if err != nil {
		panic(err)
	}

	for _, i := range productsList {
		p := i
		fmt.Printf("ID: %v\nName: %v\nPrice: %v\n", p.ID, p.Name, p.Price)
	}
}

// func main() {
// 	//sql.Open("<banco de dados>","<user>:<password>@tcp(localhost:<porta>)/<nome do DB")
// 	// Para fazer a conexão com o banco de dados MySQL.
// 	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")

// 	// // Verifica se ocorreu algum erro ao abrir a conexão com o banco de dados.
// 	if err != nil {
// 		panic(err)
// 	}

// 	// // Garante que a conexão com o banco de dados será fechada no final da função main.
// 	defer db.Close()

// 	err = deleteProduct(db, "a7bfcfba-6e94-4326-a37f-bf7de1e399df")

// }

package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Definindo a estrutura do modelo 'Product' para representar produtos.
type Product struct {
	ID           int          `gorm:"primaryKey"` // Chave primária do produto
	Name         string       // Nome do produto
	Price        float64      // Preço do produto
	CategoryID   int          // Chave estrangeira para a categoria à qual o produto pertence
	Category     Category     // Relacionamento muitos-para-um com categorias
	gorm.Model                // Modelo básico do GORM para rastrear timestamps de criação/atualização
	SerialNumber SerialNumber // Relacionamento um-para-um com números de série
}

// Definindo a estrutura do modelo 'Category' para representar categorias de produtos.
type Category struct {
	ID   int    `gorm:"primaryKey"` // Chave primária da categoria
	Code string // Código da categoria
}

// Definindo a estrutura do modelo 'SerialNumber' para representar números de série.
type SerialNumber struct {
	ID        int    `gorm:"primaryKey"` // Chave primária do número de série
	Number    string // Número de série
	ProductID int    // Chave estrangeira para o produto associado
}

func main() {
	// Abrindo uma conexão com o banco de dados MySQL usando GORM.
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	/*

		pode adicionar ao banco de dados com db.Create

	*/

	// Executando as migrações para criar tabelas no banco de dados, se ainda não existirem.
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// Consultando todos os produtos, pré-carregando informações sobre categorias e números de série.
	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)

	// Iterando sobre os produtos e exibindo informações relevantes.
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Code, product.SerialNumber.Number)
	}
}

package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Definindo a estrutura do modelo 'Category' para representar categorias de produtos.
type Category struct {
	ID       int       `gorm:"primaryKey"` // Chave primária da categoria
	Name     string    // Nome da categoria
	Products []Product // Relacionamento de um-para-muitos com produtos
}

// Definindo a estrutura do modelo 'Product' para representar produtos.
type Product struct {
	ID           int          `gorm:"primaryKey"` // Chave primária do produto
	Name         string       // Nome do produto
	Price        float64      // Preço do produto
	CategoryID   int          // Chave estrangeira para a categoria à qual o produto pertence
	Category     Category     // Relacionamento muitos-para-um com categorias
	SerialNumber SerialNumber // Relacionamento um-para-um com números de série
	gorm.Model                // Modelo básico do GORM para rastrear timestamps de criação/atualização
}

// Definindo a estrutura do modelo 'SerialNumber' para representar números de série.
type SerialNumber struct {
	ID        int    `gorm:"primaryKey"` // Chave primária do número de série
	Number    string // Número de série
	ProductID int    // Chave estrangeira para o produto associado
}

func main() {
	// Configurando a string de conexão com o banco de dados.
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	// Abrindo uma conexão com o banco de dados MySQL usando GORM.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Executando as migrações para criar tabelas no banco de dados, se ainda não existirem.
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// Criando uma nova categoria no banco de dados.
	db.Create(&Category{
		Name: "cozinha",
	})

	// Criando um novo produto associado à categoria com ID 1.
	db.Create(&Product{
		Name:       "Panela",
		Price:      49.50,
		CategoryID: 1,
	})

	// Criando um novo número de série associado ao produto com ID 1.
	db.Create(&SerialNumber{
		Number:    "192293",
		ProductID: 1,
	})

	// Consultando todas as categorias e seus produtos associados, incluindo números de série.
	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	// Iterando sobre as categorias, produtos e exibindo os números de série.
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name, "Serial Number:", product.SerialNumber.Number)
		}
	}
}

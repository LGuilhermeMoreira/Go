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
	ID         int      `gorm:"primaryKey"` // Chave primária do produto
	Name       string   // Nome do produto
	Price      float64  // Preço do produto
	CategoryID int      // Chave estrangeira para a categoria à qual o produto pertence
	Category   Category // Relacionamento muitos-para-um com categorias
	gorm.Model          // Modelo básico do GORM para rastrear timestamps de criação/atualização
}

func main() {
	// Abrindo uma conexão com o banco de dados MySQL usando GORM.
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Executando as migrações para criar tabelas no banco de dados, se ainda não existirem.
	db.AutoMigrate(&Product{}, &Category{})

	// Criando uma nova categoria no banco de dados.
	db.Create(&Category{
		Name: "Banheiro",
	})

	// Criando um novo produto associado à categoria com ID 2.
	db.Create(&Product{
		Name:       "sabonete",
		Price:      12.5,
		CategoryID: 2,
	})

	// Consultando todas as categorias e seus produtos associados.
	var categories []Category

	// Utilizando 'Preload' para carregar automaticamente os produtos relacionados.
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

	if err != nil {
		panic(err)
	}

	// Iterando sobre as categorias e exibindo seus produtos.
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println(" - ", product.Name)
		}
	}
}

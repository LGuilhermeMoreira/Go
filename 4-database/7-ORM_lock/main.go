package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Definindo a estrutura do modelo 'Category' para representar categorias de produtos.
type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"` // Relacionamento muitos-para-muitos com produtos usando a tabela de junção 'products_categories'
}

// Definindo a estrutura do modelo 'Product' para representar produtos.
type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"` // Relacionamento muitos-para-muitos com categorias usando a tabela de junção 'products_categories'
	gorm.Model
}

func main() {
	// Abrindo uma conexão com o banco de dados MySQL usando GORM.
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Executando as migrações para criar tabelas no banco de dados, se ainda não existirem.
	db.AutoMigrate(&Product{}, &Category{})

	// Iniciando uma transação.
	tx := db.Begin()

	// Criando uma instância de Category para armazenar os dados.
	var c Category

	// Selecionando e bloqueando a categoria com ID 1 para atualização.
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error

	if err != nil {
		// Em caso de erro, encerramos a transação.
		tx.Rollback()
		panic(err)
	}

	// Modificando o nome da categoria.
	c.Name = "Eletronica"

	// Salvando as alterações na transação.
	tx.Debug().Save(&c)

	// Confirmar (commit) a transação.
	tx.Commit()

	test_look
}

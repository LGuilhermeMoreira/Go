package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm: "primaryKey"`
	Name  string
	Price float64
}

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/goexpert"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	//com esse comando ele cria um proprio banco de dados baseado na struct
	//db.AutoMigrate(&Product{})

	// inserindo no banco de dados
	//db.Create(&Product{Name: "macbook m1", Price: 5000})

	// criando uma series de produtos
	/*
		products := []Products{
			{Name: "controller",Price : 500},
			{Name: "Pizza",Price : 20},
			{Name: "ipad",Price : 1450},
		}

		db.Create(&products)
	*/

	var product Product

	// buscando no banco de dados
	db.First(&product, 1) // buscando onde o id = 1
	//db.First(&product, "name = ?", "macbook m1") // buscar produtos com o nome = macbook m1

	//fmt.Printf("id: %v name: %v price: %v", product.ID, product.Name, product.Price)

	/*
		SELECIONANDO TODOS OS PRODUTOS
		var produtos []Products
		db.Find(&produtos)

		SELECIONANDO OS 2 PRIMEIROS PRODUTOS
		db.Limit(2).Find(&produtos)

		for _,i := range produtos {
			fmt.Println(i)
		}

		WHERE
		db.Where("price > ?",100).Find(&products)

		LIKE
		db.Where("name LIKE ?","%I%").Find(&produtos)
	*/

	//dando um update no banco de dados
	//db.Model(&product).Update("Name", "RTX 4090 ti") // estou trocando o valor

	/*
		OUTRA FORMA DE UPDATA
		var p Product
		db.First(&p,1)
		p.Name = "outro nome"
		db.Save(&p)

	*/

	// deletando no banco de dados
	//db.Delete(&product, 3) // deletando onde ID = 3
}

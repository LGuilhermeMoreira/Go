package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//

type Product struct {
	ID         int `gorm: "primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

type Category struct {
	ID   int `gorm : "primaryKey"`
	Code string
}

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	//com esse comando ele cria um proprio banco de dados baseado na struct
	//db.AutoMigrate(&Product{})
	db.AutoMigrate(&Product{}, &Category{})
	//=================================================================
	// db.Create(&Product{
	// 	Name:  "Alexa",
	// 	Price: 500,
	// })

	// var p Product
	// db.First(&p, 1)
	// db.Model(&p).Update("Name", "Amazon dot 4")

	//db.Delete(&Product{},[]int{3,4})
	//=================================================================
	// criando relacionamento -> produto pertence a uma categoria

	// category := Category{Code: "Eletronicos"}
	// db.Create(&category)

	// db.Create(&Product{
	// 	Name:       "Notebook",
	// 	Price:      1000,
	// 	CategoryID: category.ID,
	// })

	var produtos []Product

	// esse preload traz os dados de category para serem trabalhados
	db.Preload("Category").Find(&produtos)

	for _, i := range produtos {
		fmt.Println(i.Name, i.Category.Code)
	}

}

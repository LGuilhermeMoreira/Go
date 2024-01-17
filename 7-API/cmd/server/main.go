package main

import (
	"API/fundamentos/configs"
	"API/fundamentos/internal/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Carrega a configuração usando o caminho do arquivo .env
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open(("test.db")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

}

package main

import (
	"API/fundamentos/configs"
	"API/fundamentos/internal/entity"
	"API/fundamentos/internal/infra/database"
	"API/fundamentos/internal/infra/webserver/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Carrega a configuração usando o caminho do arquivo .env
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open(("test.db")), &gorm.Config{}) // para trocar o CGO_ENABLED: go env -w CGO_ENABLED=<0 or1>

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productHandelr := handlers.NewProductHandler(database.NewProduct(db))
	userHandler := handlers.NewUserHandler(database.NewUser(db), configs.TokenAuth, configs.JWTExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// rotas de products
	r.Post("/products", productHandelr.CreateProduct)
	r.Get("/products/{id}", productHandelr.GetProduct)
	r.Put("/products/{id}", productHandelr.UpdateProduct)
	r.Delete("/products/{id}", productHandelr.DeleteProduct)
	r.Get("/products/{sort}", productHandelr.GetAllProducts)

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	http.ListenAndServe(":8000", r)
}

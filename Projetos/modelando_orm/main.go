package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Documento struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	gorm.Model
	CPF    string
	UserID uuid.UUID `gorm:"type:uuid;index"` // Foreign key, index for performance
}

type Pedido struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	gorm.Model
	Numero uint
	UserID uuid.UUID `gorm:"type:uuid;index"` // Foreign key, index for performance
}

type Carro struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	gorm.Model
	Nome  string
	Marca string
	Users []User `gorm:"many2many:users_carros"`
}

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	gorm.Model
	Nome             string
	Documento        Documento `gorm:"foreignKey:UserID"` // HasOne relationship via Documento.UserID
	DataDeNascimento time.Time
	Pedidos          []Pedido // HasMany relationship via Pedido.UserID
	Carros           []Carro  `gorm:"many2many:users_carros"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// AutoMigrate:  Creates tables and migrates existing ones.  Good for development!
	err = db.AutoMigrate(&User{}, &Carro{}, &Documento{}, &Pedido{})
	if err != nil {
		log.Fatalf("failed to automigrate: %v", err)
	}

	// CRUD Operations for User

	// Create User
	userID := uuid.New()
	user := User{
		ID:               userID,
		Nome:             "Alice Smith",
		DataDeNascimento: time.Now().AddDate(-25, 0, 0),
	}
	result := db.Create(&user)
	if result.Error != nil {
		log.Printf("failed to create user: %v", result.Error)
	} else {
		fmt.Println("User created:", user.ID)
	}

	// Read User
	var retrievedUser User
	result = db.First(&retrievedUser, userID) // Find user by ID
	if result.Error != nil {
		log.Printf("failed to find user: %v", result.Error)
	} else {
		fmt.Println("Retrieved user:", retrievedUser.Nome)
	}

	// Update User
	retrievedUser.Nome = "Alice Updated"
	result = db.Save(&retrievedUser) // Updates all fields
	if result.Error != nil {
		log.Printf("failed to update user: %v", result.Error)
	} else {
		fmt.Println("User updated:", retrievedUser.Nome)
	}

	// Delete User
	//db.Delete(&retrievedUser) // Soft delete (sets DeletedAt)
	result = db.Unscoped().Delete(&retrievedUser) // Hard delete (permanently removes from DB)
	if result.Error != nil {
		log.Printf("failed to delete user: %v", result.Error)
	} else {
		fmt.Println("User deleted")
	}

	// CRUD Operations for Documento

	// Create Documento
	docID := uuid.New()
	doc := Documento{
		ID:     docID,
		CPF:    "111.222.333-44",
		UserID: userID, // Associate with the user we (attempted) to create above
	}
	result = db.Create(&doc)
	if result.Error != nil {
		log.Printf("failed to create document: %v", result.Error)
	} else {
		fmt.Println("Document created:", doc.ID)
	}

	// Read Documento
	var retrievedDoc Documento
	result = db.First(&retrievedDoc, docID)
	if result.Error != nil {
		log.Printf("failed to find document: %v", result.Error)
	} else {
		fmt.Println("Retrieved document:", retrievedDoc.CPF)
	}

	// Update Documento
	retrievedDoc.CPF = "555.666.777-88"
	result = db.Save(&retrievedDoc)
	if result.Error != nil {
		log.Printf("failed to update document: %v", result.Error)
	} else {
		fmt.Println("Document updated:", retrievedDoc.CPF)
	}

	// Delete Documento
	result = db.Delete(&retrievedDoc)
	if result.Error != nil {
		log.Printf("failed to delete document: %v", result.Error)
	} else {
		fmt.Println("Document deleted")
	}

	// CRUD Operations for Pedido

	// Create Pedido
	pedidoID := uuid.New()
	pedido := Pedido{
		ID:     pedidoID,
		Numero: 42,
		UserID: userID,
	}
	result = db.Create(&pedido)
	if result.Error != nil {
		log.Printf("failed to create pedido: %v", result.Error)
	} else {
		fmt.Println("Pedido created:", pedido.ID)
	}

	// Read Pedido
	var retrievedPedido Pedido
	result = db.First(&retrievedPedido, pedidoID)
	if result.Error != nil {
		log.Printf("failed to find pedido: %v", result.Error)
	} else {
		fmt.Println("Retrieved pedido:", retrievedPedido.Numero)
	}

	// Update Pedido
	retrievedPedido.Numero = 100
	result = db.Save(&retrievedPedido)
	if result.Error != nil {
		log.Printf("failed to update pedido: %v", result.Error)
	} else {
		fmt.Println("Pedido updated:", retrievedPedido.Numero)
	}

	// Delete Pedido
	result = db.Delete(&retrievedPedido)
	if result.Error != nil {
		log.Printf("failed to delete pedido: %v", result.Error)
	} else {
		fmt.Println("Pedido deleted")
	}

	// CRUD Operations for Carro

	// Create Carro
	carroID := uuid.New()
	carro := Carro{
		ID:    carroID,
		Nome:  "Fusca",
		Marca: "Volkswagen",
	}
	result = db.Create(&carro)
	if result.Error != nil {
		log.Printf("failed to create carro: %v", result.Error)
	} else {
		fmt.Println("Carro created:", carro.ID)
	}

	// Read Carro
	var retrievedCarro Carro
	result = db.First(&retrievedCarro, carroID)
	if result.Error != nil {
		log.Printf("failed to find carro: %v", result.Error)
	} else {
		fmt.Println("Retrieved carro:", retrievedCarro.Nome)
	}

	// Update Carro
	retrievedCarro.Nome = "Novo Fusca"
	result = db.Save(&retrievedCarro)
	if result.Error != nil {
		log.Printf("failed to update carro: %v", result.Error)
	} else {
		fmt.Println("Carro updated:", retrievedCarro.Nome)
	}

	// Delete Carro
	result = db.Delete(&retrievedCarro)
	if result.Error != nil {
		log.Printf("failed to delete carro: %v", result.Error)
	} else {
		fmt.Println("Carro deleted")
	}

	// Example of adding a car to a user (many-to-many)

	var userToUpdate User

	result = db.Where("nome = ?", "Alice Smith").First(&userToUpdate)

	if result.Error != nil {
		log.Printf("failed to find Alice: %v", result.Error)
	}

	var carroToAdd Carro

	result = db.Where("nome = ?", "Novo Fusca").First(&carroToAdd)

	if result.Error != nil {
		log.Printf("failed to find Novo Fusca: %v", result.Error)
	}

	db.Model(&userToUpdate).Association("Carros").Append(&carroToAdd)

}

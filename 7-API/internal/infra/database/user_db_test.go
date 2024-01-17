package database

import (
	"API/fundamentos/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	// Substitua os valores de DSN abaixo com as informações do seu banco de dados MySQL
	dsn := "root:root@tcp(localhost:3305)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("gui", "guilhermemoreira@alu.ufc.br", "leite123")

	userDb := NewUser(db)

	err = userDb.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestFindUserByEmail(t *testing.T) {
	dsn := "root:root@tcp(localhost:3305)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("guilherme", "mail.com", "123444")

	userDb := NewUser(db)

	err = userDb.Create(user)

	assert.Nil(t, err)

	userFound, err := userDb.FindByEmail(user.Email)

	assert.Nil(t, err)

	assert.Equal(t, user.ID, userFound.ID)

	assert.Equal(t, user.Name, userFound.Name)

	assert.Equal(t, user.Email, userFound.Email)

	assert.NotNil(t, userFound.Password)
}

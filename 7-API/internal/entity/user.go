package entity

import (
	"API/fundamentos/pkg/entity"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"` // password nunca sera exibido
}

func NewUser(name, email, password string) (*User, error) {
	// nunca é guardado a senha direto no banco

	// criando um hash dessa senha
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) Show() {
	fmt.Println(u.Name, u.Email, u.Password)
}

// função que comprar
func (u *User) ValidatePassword(password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}

package entity

import (
	"API/fundamentos/pkg/entity"
	"errors"
	"time"
)

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrIDIsInvalid     = errors.New("id is invalid")
	ErrPriceIsRequired = errors.New("price is required")
	ErrPriceIsInvalid  = errors.New("price is invalid")
	ErrNameIsRequired  = errors.New("name is required")
)

type Product struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	CreateAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := Product{
		Name:     name,
		Price:    price,
		CreateAt: time.Now(),
		ID:       entity.NewID(),
	}

	err := product.Validate()

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrIDIsInvalid
	}

	if p.Price == 0 {
		return ErrPriceIsRequired
	}

	if p.Price < 0 {
		return ErrPriceIsInvalid
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	return nil
}

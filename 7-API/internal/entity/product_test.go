package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("pia", 14.0)

	assert.Nil(t, err)

	assert.NotNil(t, p)

	assert.NotEmpty(t, p.ID)

	assert.Equal(t, p.Name, "pia")

	assert.Equal(t, p.Price, 14.0)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("pia", 14.0)
	assert.Nil(t, p.Validate())
	assert.Nil(t, err)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 14.0)

	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("pia", -14.0)

	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceIsInvalid, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("pia", 0.0)

	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceIsRequired, err)
}

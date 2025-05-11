package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Nintendo Switch 2", 4500.00)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.Name)
	assert.NotEmpty(t, product.Price)
	assert.NotEmpty(t, product.CreatedAt)

	assert.Equal(t, "Nintendo Switch 2", product.Name)
	assert.Equal(t, 4500.00, product.Price)
	assert.Nil(t, product.Validate())
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 4500.00)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNameIsRequired, err)
	assert.Nil(t, product)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Nintendo Switch 2", 0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceIsRequired, err)
	assert.Nil(t, product)
}

func TestProductWhenPriceIsInvalide(t *testing.T) {
	product, err := NewProduct("Nintendo Switch 2", -1)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceIsRequired, err)
	assert.Nil(t, product)
}
package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Test Product", 100)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Test Product", product.Name)
	assert.Equal(t, 100, product.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	_, err := NewProduct("", 100)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	_, err := NewProduct("Test Product", 0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	_, err := NewProduct("Test Product", -100)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
}
func TestProductIsValidate(t *testing.T) {
	product, err := NewProduct("Test Product", 100)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validate())
}

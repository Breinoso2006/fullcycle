package database

import (
	"fmt"
	"math/rand/v2"
	"testing"
	"time"

	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/08-apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Nintendo Switch 2", 10.0)
	assert.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEqual(t, "", product.ID.String())
}

func TestFindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 25; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Nintendo Switch %d", i), rand.Float64())
		assert.NoError(t, err)
		db.Create(product)
		time.Sleep(time.Millisecond * 5)
	}
	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Nintendo Switch 1", products[0].Name)
	assert.Equal(t, "Nintendo Switch 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Nintendo Switch 11", products[0].Name)
	assert.Equal(t, "Nintendo Switch 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 4)
	assert.Equal(t, "Nintendo Switch 21", products[0].Name)
	assert.Equal(t, "Nintendo Switch 24", products[3].Name)
}

func TestFindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Nintendo Switch 1", 10.0)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.Name, productFound.Name)

}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Nintendo Switch 1", 10.0)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	productFound.Name = "Nintendo Switch 2"
	err = productDB.Update(productFound)
	assert.NoError(t, err)

	productFoundUpdated, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Nintendo Switch 2", productFoundUpdated.Name)

}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Nintendo Switch 1", 10.0)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.Error(t, err)
	assert.Nil(t, productFound)

}

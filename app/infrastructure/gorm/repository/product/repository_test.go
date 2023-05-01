package product_test

import (
	"log"
	"os"
	"queryservice/domain/repository"
	"queryservice/infrastructure/gorm/handler"
	"queryservice/infrastructure/gorm/repository/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rep repository.ProductRepository

func TestMain(m *testing.M) {
	db, err := handler.NewGORMConnector().Open()
	if err != nil {
		os.Exit(-1)
		return
	}
	rep, err = product.NewProductRepositoryGorm(db)
	if err != nil {
		os.Exit(-1)
		return
	}
	// テストの実行
	status := m.Run()
	os.Exit(status)
}

func TestFindAll(t *testing.T) {
	products, err := rep.FindAll()
	if err != nil {
		assert.Error(t, err)
		return
	}
	for _, product := range products {
		log.Println(product)
	}
	assert.True(t, len(products) > 0)
}

func TestFindByIdOK(t *testing.T) {
	product, err := rep.FindById("e4850253-f363-4e79-8110-7335e4af45be")
	if err != nil {
		assert.Error(t, err)
		return
	}
	log.Println(product)
	assert.NotNil(t, product)
}
func TestFindByIdNG(t *testing.T) {
	product, err := rep.FindById("e4850253-f363-4e79-8110-7335e4af45b0")
	if err != nil {
		assert.Error(t, err)
		return
	}

	assert.True(t, product.Id == "")
	assert.True(t, product.Name == "")
	assert.True(t, product.Price == 0)
}

func TestFindByKeywordOK(t *testing.T) {
	products, err := rep.FindByKeyword("ペン")
	if err != nil {
		assert.Error(t, err)
		return
	}
	for _, product := range products {
		log.Println(product)
	}
	assert.True(t, len(products) > 0)
}
func TestFindByKeywordNG(t *testing.T) {
	products, err := rep.FindByKeyword("おお")
	if err != nil {
		assert.Error(t, err)
		return
	}
	assert.True(t, len(products) == 0)
}

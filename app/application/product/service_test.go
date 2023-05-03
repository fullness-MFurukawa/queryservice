package product_test

import (
	"context"
	"log"
	"os"
	"queryservice/application/pb"
	"queryservice/application/product"
	"queryservice/infrastructure/gorm/provider"
	"testing"

	"github.com/stretchr/testify/assert"
)

var service *product.ProductService

func TestMain(m *testing.M) {
	provider, err := provider.NewRepositoryProvider()
	if err != nil {
		panic(err)
	}
	service = product.NewProductService(provider.ProductRep)

	// テストの実行
	status := m.Run()
	os.Exit(status)
}

func TestList(t *testing.T) {
	param := pb.ProductParam{}
	resp, err := service.List(context.Background(), &param)
	if err != nil {
		assert.Error(t, err)
		return
	}
	for _, product := range resp.GetProducts() {
		log.Println(product)
	}
	assert.True(t, len(resp.GetProducts()) > 0)
}

func TestByIdOK(t *testing.T) {
	id := "8f81a72a-58ef-422b-b472-d982e8665292"
	param := pb.ProductParam{Id: &id}
	resp, err := service.ById(context.Background(), &param)
	if err != nil {
		assert.Error(t, err)
	}
	log.Println(resp.Result)
	assert.NotNil(t, resp.Result)
}
func TestByIdNG(t *testing.T) {
	id := "8f81a72a-58ef-422b-b472-d982e8665290"
	param := pb.ProductParam{Id: &id}
	resp, err := service.ById(context.Background(), &param)
	if err != nil {
		assert.Error(t, err)
	}
	log.Println(resp.Result)
	assert.NotNil(t, resp.Result)
}

func TestByKeywordOK(t *testing.T) {
	keyword := "ペン"
	param := pb.ProductParam{Keyword: &keyword}
	resp, err := service.ByKeyword(context.Background(), &param)
	if err != nil {
		assert.Error(t, err)
	}
	for _, product := range resp.GetProducts() {
		log.Println(product)
	}
	assert.True(t, len(resp.GetProducts()) > 0)
}

func TestByKeywordNG(t *testing.T) {
	keyword := "ABC"
	param := pb.ProductParam{Keyword: &keyword}
	resp, err := service.ByKeyword(context.Background(), &param)
	if err != nil {
		assert.Error(t, err)
	}
	log.Println(resp.GetError())
	assert.NotNil(t, resp.GetError())
}

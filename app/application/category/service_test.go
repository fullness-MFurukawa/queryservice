package category_test

import (
	"context"
	"log"
	"os"
	"queryservice/application/category"
	"queryservice/application/pb"
	"queryservice/infrastructure/gorm/provider"
	"testing"

	"github.com/stretchr/testify/assert"
)

var service *category.CategoryService

func TestMain(m *testing.M) {
	provider, err := provider.NewRepositoryProvider()
	if err != nil {
		panic(err)
	}
	service = category.NewCategoryService(provider.CategoryRep)

	// テストの実行
	status := m.Run()
	os.Exit(status)
}

func TestList(t *testing.T) {
	param := pb.CategoryParam{}
	resp, err := service.List(context.Background(), &param)
	if err != nil {
		assert.Error(t, err)
		return
	}
	for _, category := range resp.GetCategories() {
		log.Println(category)
	}
	assert.True(t, len(resp.GetCategories()) > 0)
}

func TestByIdOK(t *testing.T) {
	id := "762bd1ea-9700-4bab-a28d-6cbebf20ddc2"
	param := pb.CategoryParam{Id: &id}
	resp, err := service.ById(context.Background(), &param)
	if err != nil {
		assert.Error(t, err)
	}
	log.Println(resp.Result)
	assert.NotNil(t, resp.Result)
}

func TestByIdNG(t *testing.T) {
	id := "762bd1ea-9700-4bab-a28d-6cbebf20ddc0"
	param := pb.CategoryParam{Id: &id}
	resp, err := service.ById(context.Background(), &param)
	if err != nil {
		assert.Error(t, err)
	}
	log.Println(resp.Result)
	assert.NotNil(t, resp.Result)
}

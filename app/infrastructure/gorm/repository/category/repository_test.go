package category_test

import (
	"log"
	"os"
	"queryservice/domain/repository"
	"queryservice/infrastructure/gorm/handler"
	"queryservice/infrastructure/gorm/repository/category"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rep repository.CategoryRepository

func TestMain(m *testing.M) {
	db, err := handler.NewGORMConnector().Open()
	if err != nil {
		os.Exit(-1)
		return
	}
	rep, err = category.NewCategoryRepositoryGorm(db)
	if err != nil {
		os.Exit(-1)
		return
	}
	// テストの実行
	status := m.Run()
	os.Exit(status)
}

func TestCategoryFindAll(t *testing.T) {
	categories, err := rep.FindAll()
	if err != nil {
		assert.Error(t, err)
		return
	}
	for _, category := range categories {
		log.Println(category)
	}
	assert.True(t, len(categories) > 0)
}
func TestFindByIdOk(t *testing.T) {
	category, err := rep.FindById("b1524011-b6af-417e-8bf2-f449dd58b5c0")
	if err != nil {
		assert.Error(t, err)
		return
	}
	log.Println(category)
	assert.NotNil(t, category)
}

func TestFindByIdNG(t *testing.T) {
	category, err := rep.FindById("b1524011-b6af-417e-8bf2-f449dd58b5c1")
	if err != nil {
		assert.Error(t, err)
		return
	}
	assert.True(t, category.Id == "")
	assert.True(t, category.Name == "")
}

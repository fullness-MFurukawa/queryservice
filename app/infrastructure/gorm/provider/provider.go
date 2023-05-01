package provider

import (
	"queryservice/domain/repository"
	"queryservice/infrastructure/gorm/handler"
	"queryservice/infrastructure/gorm/repository/category"
	"queryservice/infrastructure/gorm/repository/product"
)

// *****
// リポジトリの提供
// *****
type RepositoryProvider struct {
	CategoryRep repository.CategoryRepository
	ProductRep  repository.ProductRepository
}

/*
コンストラクタ
*/
func NewRepositoryProvider() (*RepositoryProvider, error) {
	db, err := handler.NewGORMConnector().Open()
	if err != nil {
		return nil, err
	}
	categoryRep, err := category.NewCategoryRepositoryGorm(db)
	if err != nil {
		return nil, err
	}
	productRep, err := product.NewProductRepositoryGorm(db)
	if err != nil {
		return nil, err
	}
	return &RepositoryProvider{CategoryRep: categoryRep, ProductRep: productRep}, nil
}

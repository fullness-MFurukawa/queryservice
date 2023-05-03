package provider

import (
	"queryservice/application/category"
	"queryservice/application/product"
	"queryservice/infrastructure/gorm/provider"
)

type ServiceProvider struct {
	CategoryService category.CategoryService
	ProductService  product.ProductService
}

/*
コンストラクタ
*/
func NewServiceProvider() *ServiceProvider {
	provider, err := provider.NewRepositoryProvider()
	if err != nil {
		panic(err)
	}
	return &ServiceProvider{
		CategoryService: *category.NewCategoryService(provider.CategoryRep),
		ProductService:  *product.NewProductService(provider.ProductRep),
	}
}

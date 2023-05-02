package privider

import (
	cs "queryservice/application/category"
	ps "queryservice/application/product"
)

type ServiceProvider struct {
	CategoryService cs.CategoryService
	ProductService  ps.ProductService
}

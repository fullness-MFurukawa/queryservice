package product

import (
	"errors"
	"queryservice/domain/model"
	"queryservice/domain/repository"

	"gorm.io/gorm"
)

// *****
// GORMを利用した商品リポジトリ
// *****
type ProductRepositoryGorm struct {
	db *gorm.DB
}

/*
すべての商品を取得する
*/
func (rep *ProductRepositoryGorm) FindAll() ([]model.Product, error) {
	var products = []model.Product{}
	if result := rep.db.Table("product").Find(&products); result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

/*
引数指定された商品Idの商品を取得する
*/
func (rep *ProductRepositoryGorm) FindById(id string) (*model.Product, error) {
	var product model.Product
	if result := rep.db.Table("product").Where("obj_id = ?", id).Find(&product); result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

/*
引数指定されたキーワードを含む商品を取得する
*/
func (rep *ProductRepositoryGorm) FindByKeyword(keyword string) ([]model.Product, error) {
	var products = []model.Product{}
	if result := rep.db.Table("product").Where("name LIKE ?", "%"+keyword+"%").Find(&products); result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

/*
コンストラクタ
*/
func NewProductRepositoryGorm(db *gorm.DB) (repository.ProductRepository, error) {
	if db == nil {
		return nil, errors.New("データベース接続がnilです。")
	}
	return &ProductRepositoryGorm{db: db}, nil
}

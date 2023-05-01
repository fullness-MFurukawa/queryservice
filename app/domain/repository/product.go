package repository

import "queryservice/domain/model"

// *****
// 商品ポジトリインターフェース
// *****
type ProductRepository interface {
	// すべての商品を取得する
	FindAll() ([]model.Product, error)
	// 引数指定された商品Idの商品を取得する
	FindById(id string) (*model.Product, error)
	// 引数指定されたキーワードを含む商品を取得する
	FindByKeyword(keyword string) ([]model.Product, error)
}

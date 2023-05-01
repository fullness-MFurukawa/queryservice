package repository

import "queryservice/domain/model"

// *****
// 商品カテゴリリポジトリインターフェース
// *****
type CategoryRepository interface {
	// すべての商品カテゴリを取得する
	FindAll() ([]model.Category, error)
	// 引数指定された商品カテゴリIdのカテゴリを取得する
	FindById(id string) (*model.Category, error)
}

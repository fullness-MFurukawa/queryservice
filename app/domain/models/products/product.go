package products

import "queryservice/domain/models/categories"

// 商品エンティティ
type Product struct {
	id       string               // 商品番号
	name     string               // 商品名名
	price    uint32               // 単価
	category *categories.Category // カテゴリ
}

// コンストラクタ
func NewProduct(id string, name string, price uint32, category *categories.Category) *Product {
	return &Product{id: id, name: name, price: price, category: category}
}

// ゲッター
func (ins *Product) Id() string {
	return ins.id
}
func (ins *Product) Name() string {
	return ins.name
}
func (ins *Product) Price() uint32 {
	return ins.price
}
func (ins *Product) Category() *categories.Category {
	return ins.category
}

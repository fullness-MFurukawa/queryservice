package products

// 商品変換インターフェイス
type ProductAdapter interface {
	// Productから他のモデルに変換
	Convert(source *Product) any
	// 他のモデルからProductに変換
	ReBuild(source any) (dest *Product, err error)
}

package products

import "context"

// 商品検索リポジトリインターフェイス
type ProductRepository interface {
	// 商品リストを取得する
	List(ctx context.Context) ([]*Product, error)
	// 指定された商品IDの商品を取得する
	FindByProductId(ctx context.Context, productid string) (*Product, error)
	// 指定されたキーワードの商品を取得する
	FindByProductNameLike(ctx context.Context, keyword string) ([]*Product, error)
}

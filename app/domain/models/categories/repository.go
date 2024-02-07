package categories

import "context"

// カテゴリ検索リポジトリインターフェイス
type CategoryRepository interface {
	// カテゴリリストを取得する
	List(ctx context.Context) ([]*Category, error)
	// 指定されたカテゴリIDのカテゴリを取得する
	FindByCategoryId(ctx context.Context, categoryid string) (*Category, error)
}

package categories

// カテゴリ変換インターフェイス
type CategoryAdapter interface {
	// Categoryから他のモデルに変換
	Convert(source *Category) any
	// 他のモデルからCategoryに変換
	ReBuild(source any) (dest *Category, err error)
}

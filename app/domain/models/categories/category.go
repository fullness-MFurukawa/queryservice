package categories

// カテゴリエンティティ
type Category struct {
	id   string
	name string
}

// コンストラクタ
func NewCategory(id string, name string) *Category {
	return &Category{id: id, name: name}
}

// ゲッター
func (ins *Category) Id() string {
	return ins.id
}
func (ins *Category) Name() string {
	return ins.name
}

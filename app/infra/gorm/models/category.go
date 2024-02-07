package models

// 商品カテゴリ
type Category struct {
	// id(主キー)
	ID int `gorm:"column:c_key;primaryKeY"`
	// カテゴリ番号
	ObjId string `gorm:"column:c_id;unique"`
	// カテゴリ名
	Name string `gorm:"column:c_name"`
}

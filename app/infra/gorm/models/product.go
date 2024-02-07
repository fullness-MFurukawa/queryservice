package models

// 商品データの格納
type Product struct {
	ID           int    `gorm:"column:p_key;primaryKey"` // id(主キー)
	ObjId        string `gorm:"column:p_id;unique"`      // 商品番号
	Name         string `gorm:"column:p_name"`           // 商品名名
	Price        uint32 `gorm:"column:p_price"`          //	単価
	CategoryId   string `gorm:"column:c_id"`             // カテゴリID
	CategoryName string `gorm:"column:c_name"`           // カテゴリ
}

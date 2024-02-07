package adapter

import (
	"queryservice/domain/models/categories"
	"queryservice/errs"
	"queryservice/infra/gorm/models"
)

// カテゴリ変換インターフェイスの実装
type categoryAdapterImpl struct{}

// コンストラクタ
func NewcategoryAdapterImpl() categories.CategoryAdapter {
	return &categoryAdapterImpl{}
}

// Categoryから他のモデルに変換
func (ins *categoryAdapterImpl) Convert(source *categories.Category) any {
	return &models.Category{
		ObjId: source.Id(),
		Name:  source.Name(),
	}
}

// 他のモデルからCategoryに変換
func (ins *categoryAdapterImpl) ReBuild(source any) (dest *categories.Category, err error) {
	if c, ok := source.(*models.Category); ok {
		dest = categories.NewCategory(c.ObjId, c.Name)
	} else {
		err = errs.NewInternalError("*models.Category以外の値が指定されました。")
	}
	return
}

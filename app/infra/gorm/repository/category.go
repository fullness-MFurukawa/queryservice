package repository

import (
	"context"
	"fmt"
	"queryservice/domain/models/categories"
	"queryservice/errs"
	"queryservice/infra/gorm/models"

	"gorm.io/gorm"
)

const (
	CATEGORY_TABLE   = "category"
	CATEGORY_COLUMNS = "id AS c_key,obj_id AS c_id,name AS c_name"
	CATEGORY_WHERE   = "obj_id = ?"
)

// カテゴリ検索リポジトリインターフェイスの実装
type categoryRepositoryGORM struct {
	db      *gorm.DB
	adapter categories.CategoryAdapter
}

// コンストラクタ
func NewcategoryRepositoryGORM(db *gorm.DB, adapter categories.CategoryAdapter) categories.CategoryRepository {
	return &categoryRepositoryGORM{db: db, adapter: adapter}
}

// カテゴリリストを取得する
func (ins *categoryRepositoryGORM) List(ctx context.Context) ([]*categories.Category, error) {
	var models = []*models.Category{}
	if result := ins.db.WithContext(ctx).
		Table(CATEGORY_TABLE).Select(CATEGORY_COLUMNS).Find(&models); result.Error != nil {
		return nil, errs.NewCRUDError(result.Error.Error())
	}

	var categories []*categories.Category
	for _, model := range models {
		if category, err := ins.adapter.ReBuild(model); err != nil {
			return nil, err
		} else {
			categories = append(categories, category)
		}
	}
	return categories, nil
}

// 指定されたカテゴリIDのカテゴリを取得する
func (ins *categoryRepositoryGORM) FindByCategoryId(ctx context.Context, categoryid string) (*categories.Category, error) {
	var model *models.Category
	if result := ins.db.WithContext(ctx).Table(CATEGORY_TABLE).
		Select(CATEGORY_COLUMNS).Where(CATEGORY_WHERE, categoryid).Find(&model); result.Error != nil {
		return nil, result.Error
	}
	if model.ID == 0 { // レコードが存在しない
		return nil, errs.NewCRUDError(fmt.Sprintf("カテゴリID:%sは存在しません。", categoryid))
	}
	if category, err := ins.adapter.ReBuild(model); err != nil {
		return nil, err
	} else {
		return category, nil
	}
}

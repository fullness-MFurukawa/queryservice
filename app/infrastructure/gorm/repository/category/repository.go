package category

import (
	"errors"
	"queryservice/apperror"
	"queryservice/domain/model"
	"queryservice/domain/repository"

	"gorm.io/gorm"
)

// *****
// GORMを利用した商品カテゴリリポジトリ
// *****
type CategoryRepositoryGorm struct {
	db *gorm.DB
}

/*
すべての商品カテゴリを取得する
*/
func (rep *CategoryRepositoryGorm) FindAll() ([]model.Category, error) {
	var categories = []model.Category{}
	// すべての商品カテゴリを取得する
	if result := rep.db.Table("category").Find(&categories); result.Error != nil {
		return nil, apperror.NewInternalError(result.Error)
	}
	return categories, nil
}

/*
引数指定された商品カテゴリIdのカテゴリを取得する
*/
func (rep *CategoryRepositoryGorm) FindById(id string) (*model.Category, error) {
	var category model.Category
	if result := rep.db.Table("category").Where("obj_id = ?", id).Find(&category); result.Error != nil {
		return nil, apperror.NewInternalError(result.Error)
	}
	return &category, nil
}

/*
コンストラクタ
*/
func NewCategoryRepositoryGorm(db *gorm.DB) (repository.CategoryRepository, error) {
	if db == nil {
		return nil, errors.New("データベース接続がnilです。")
	}
	return &CategoryRepositoryGorm{db: db}, nil
}

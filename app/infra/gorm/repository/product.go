package repository

import (
	"context"
	"fmt"
	"queryservice/domain/models/products"
	"queryservice/errs"
	"queryservice/infra/gorm/handler"
	"queryservice/infra/gorm/models"

	"gorm.io/gorm"
)

// テーブル名とSQLステートメント
const (
	PRODUCT_TABLE  = "product"
	PRODUCT_COLUMS = "product.obj_id AS p_id , product.name AS p_name , product.price AS p_price , product.category_id AS c_id , category.name AS c_name"
	PRODUCT_JOIN   = "JOIN category ON product.category_id = category.obj_id"
	PRODUCT_WHERE  = "product.obj_id = ?"
	PRODUCT_LIKE   = "product.name LIKE ?"
)

// 商品検索リポジトリインターフェイスの実装
type productRepositoryGORM struct {
	db      *gorm.DB
	adapter products.ProductAdapter
}

// コンストラクタ
func NewproductRepositoryGORM(db *gorm.DB, adapter products.ProductAdapter) products.ProductRepository {
	return &productRepositoryGORM{db: db, adapter: adapter}
}

// 商品リストを取得する
func (ins *productRepositoryGORM) List(ctx context.Context) ([]*products.Product, error) {
	models := []*models.Product{}
	if result := ins.db.WithContext(ctx). // Contextを設定する
						Table(PRODUCT_TABLE).                // アクセスするテーブル名を設定する
						Select(PRODUCT_COLUMS).              // 取得する列を設定する
						Joins(PRODUCT_JOIN).                 // カテゴリを結合する
						Find(&models); result.Error != nil { // 問合せした結果をスライスに格納する
		return nil, handler.DBErrHandler(result.Error)
	}
	if products, err := ins.createSlice(models); err != nil {
		return nil, err
	} else {
		return products, nil
	}
}

// 指定された商品IDの商品を取得する
func (ins *productRepositoryGORM) FindByProductId(ctx context.Context, productid string) (*products.Product, error) {
	model := models.Product{}
	if result := ins.db.WithContext(ctx). // Contextを設定する
						Table(PRODUCT_TABLE).               // アクセスするテーブル名を設定する
						Select(PRODUCT_COLUMS).             // 取得する列を設定する
						Joins(PRODUCT_JOIN).                // カテゴリを結合する
						Where(PRODUCT_WHERE, productid).    // 問合せ条件と値を設定する
						Find(&model); result.Error != nil { // 問合せした結果をスライスに格納する
		return nil, handler.DBErrHandler(result.Error)
	}
	if model.ObjId == "" { //コードが存在しない
		return nil, errs.NewCRUDError(fmt.Sprintf("商品ID:%sは存在しません。", productid))
	}
	if product, err := ins.adapter.ReBuild(model); err != nil {
		return nil, err
	} else {
		return product, nil
	}
}

// 指定されたキーワードの商品を取得する
func (ins *productRepositoryGORM) FindByProductNameLike(ctx context.Context, keyword string) ([]*products.Product, error) {
	models := []*models.Product{}
	if result := ins.db.WithContext(ctx). // Contextを設定する
						Table(PRODUCT_TABLE).                 // アクセスするテーブル名を設定する
						Select(PRODUCT_COLUMS).               // カテゴリを結合する
						Joins(PRODUCT_JOIN).                  // カテゴリを結合する
						Where(PRODUCT_LIKE, "%"+keyword+"%"). // 問合せ条件と値を設定する
						Find(&models); result.Error != nil {  // 問合せした結果をスライスに格納する
		return nil, handler.DBErrHandler(result.Error)
	}
	if len(models) == 0 { // レコードが存在しない
		return nil, errs.NewCRUDError(fmt.Sprintf("[%s]を含む商品は存在しません。", keyword))
	}
	if products, err := ins.createSlice(models); err != nil {
		return nil, err
	} else {
		return products, nil
	}
}

// 問合せ結果からエンティティのスライスを生成する
func (ins *productRepositoryGORM) createSlice(results []*models.Product) ([]*products.Product, error) {
	var products []*products.Product
	for _, result := range results {
		product, err := ins.adapter.ReBuild(result)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

package builder

import (
	"github.com/fullness-MFurukawa/samplepb/pb"
)

// 実行結果をXXXResult型に変換するインターフェイス
type ResultBuilder interface {
	// *categories.Categoryを*pb.CategtoryResultに変換する
	BuildCategoryResult(source any) *pb.CategoryResult
	// []*categories.Categoryを*pb.CategoriesResultに変換する
	BuildCategoriesResult(source any) *pb.CategoriesResult
	// *products.Productを*pbProductResultに変換する
	BuildProductResult(source any) *pb.ProductResult
	// []*product.Productを*pb.ProsuctsResultに変換する
	BuildProductsResult(source any) *pb.ProductsResult
	// errs.CRUDError、errs.InternalErrorを*pb.Errorに変換する
	BuildErrorResult(source any) *pb.Error
}

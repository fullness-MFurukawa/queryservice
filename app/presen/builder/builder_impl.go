package builder

import (
	"log"
	"queryservice/domain/models/categories"
	"queryservice/domain/models/products"
	"queryservice/errs"

	"github.com/fullness-MFurukawa/samplepb/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type resultBuilderImpl struct{}

func NewresultBuilderImpl() ResultBuilder {
	return &resultBuilderImpl{}
}

// *categories.Categoryを*pb.CategtoryResultに変換する
func (ins *resultBuilderImpl) BuildCategoryResult(source any) *pb.CategoryResult {
	// CategoryResultを生成する
	result := &pb.CategoryResult{Timestamp: timestamppb.Now()}
	// *categories.Categoryであるかを検証する
	if category, ok := source.(*categories.Category); ok {
		// Resultフィールドに問合せ結果を設定する
		result.Result = &pb.CategoryResult_Category{
			Category: &pb.Category{Id: category.Id(), Name: category.Name()},
		}
	} else {
		// Resultフィールドにエラーを設定する
		result.Result = &pb.CategoryResult_Error{Error: ins.BuildErrorResult(source)}
	}
	return result
}

// []*categories.Categoryを*pb.CategoriesResultに変換する
func (ins *resultBuilderImpl) BuildCategoriesResult(source any) *pb.CategoriesResult {
	// CategoriesResultを生成する
	result := &pb.CategoriesResult{Timestamp: timestamppb.Now()}
	// []categories.Category型であるかを検証する
	if categories, ok := source.([]*categories.Category); ok {
		// 問合せ結果を設定する
		c := []*pb.Category{}
		for _, category := range categories {
			c = append(c, &pb.Category{Id: category.Id(), Name: category.Name()})
		}
		result.Categories = c
	} else {
		// Errorフィールドにエラーを設定する
		result.Error = ins.BuildErrorResult(source)
	}
	return result
}

// *products.Productを*pbProductResultに変換する
func (ins *resultBuilderImpl) BuildProductResult(source any) *pb.ProductResult {
	// ProductResult型を生成する
	result := &pb.ProductResult{Timestamp: timestamppb.Now()}
	// *products.Productであるかを検証する
	if product, ok := source.(*products.Product); ok {
		// Resultフィールドに問合せ結果を設定する
		c := &pb.Category{Id: product.Id(), Name: product.Name()}
		result.Result = &pb.ProductResult_Product{
			Product: &pb.Product{Id: product.Id(), Name: product.Name(), Price: int32(product.Price()), Category: c},
		}
	} else {
		// Resultフィールドにエラーを設定する
		result.Result = &pb.ProductResult_Error{Error: ins.BuildErrorResult(source)}
	}
	return result
}

// []*product.Productを*pb.ProsuctsResultに変換する
func (ins *resultBuilderImpl) BuildProductsResult(source any) *pb.ProductsResult {
	// ProductsResult型を生成する
	result := &pb.ProductsResult{Timestamp: timestamppb.Now()}
	// []*products.Product型であるかを検証する
	if products, ok := source.([]*products.Product); ok {
		p := []*pb.Product{} // 問合せ結果を設定する
		for _, product := range products {
			c := &pb.Category{Id: product.Category().Id(), Name: product.Category().Name()}
			p = append(p, &pb.Product{Id: product.Id(), Name: product.Name(), Price: int32(product.Price()), Category: c})
		}
		result.Products = p
	} else {
		// Errorフィールドにエラーを設定する
		result.Error = ins.BuildErrorResult(source)
	}
	return result
}

// errs.CRUDError、errs.InternalErrorを*pb.Errorに変換する
func (ins *resultBuilderImpl) BuildErrorResult(source any) *pb.Error {
	switch v := source.(type) {
	case *errs.CRUDError:
		return &pb.Error{Type: "CRUD Error", Message: v.Error()}
	case *errs.InternalError:
		return &pb.Error{Type: "INTERNAL Error", Message: "只今、サービスを提供できません。"}
	default:
		log.Println("対応できないエラー型が指定されました。")
		return &pb.Error{Type: "INTERNAL Error", Message: "只今、サービスを提供できません。"}
	}
}

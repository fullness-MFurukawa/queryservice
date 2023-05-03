package product

import (
	"context"
	"fmt"
	"queryservice/apperror"
	"queryservice/application/pb"
	"queryservice/domain/repository"
)

// *****
// 　商品サービス
// *****
type ProductService struct {
	rep repository.ProductRepository
	// 生成したUnimplementedCategtoryServiceServerを埋め込む
	pb.UnimplementedProductServiceServer
}

/*
すべての商品を取得して返す
*/
func (svr *ProductService) List(ctx context.Context, param *pb.ProductParam) (*pb.ProductsResult, error) {
	products, err := svr.rep.FindAll()
	if err != nil {
		return &pb.ProductsResult{Error: &pb.Error{Message: err.Error()}, Products: nil}, nil
	}
	results := make([]*pb.Product, 0, len(products))
	for _, product := range products {
		results = append(results, &pb.Product{Id: product.Id, Name: product.Name, Price: int32(product.Price)})
	}
	return &pb.ProductsResult{Products: results, Error: nil}, nil
}

/*
指定されたIDの商品を取得して返す
*/
func (svr *ProductService) ById(ctx context.Context, param *pb.ProductParam) (*pb.ProductResult, error) {
	product, err := svr.rep.FindById(param.GetId())
	if err != nil {
		return &pb.ProductResult{Result: &pb.ProductResult_Error{Error: &pb.Error{Message: err.Error()}}}, nil
	}
	if product.Id == "" && product.Name == "" && product.Price == 0 {
		// エラーを生成する
		err := apperror.NewAppError("404", fmt.Sprintf("id[%s]に該当する商品が見つかりませんでした。", param.GetId()))
		return &pb.ProductResult{Result: &pb.ProductResult_Error{Error: &pb.Error{Message: err.Error()}}}, nil
	}
	result := pb.Product{Id: product.Id, Name: product.Name, Price: int32(product.Price)}
	return &pb.ProductResult{Result: &pb.ProductResult_Product{Product: &result}}, nil
}

/*
指定されたキーワードの商品を取得して返す
*/
func (svr *ProductService) ByKeyword(ctx context.Context, param *pb.ProductParam) (*pb.ProductsResult, error) {
	products, err := svr.rep.FindByKeyword(param.GetKeyword())
	if err != nil {
		return &pb.ProductsResult{Error: &pb.Error{Message: err.Error()}, Products: nil}, nil
	}
	if len(products) == 0 {
		// エラーを生成する
		err := apperror.NewAppError("404", fmt.Sprintf("[%s]を含む商品は見つかりませんでした。", param.GetKeyword()))
		return &pb.ProductsResult{Error: &pb.Error{Message: err.Error()}, Products: nil}, nil
	}
	results := make([]*pb.Product, 0, len(products))
	for _, product := range products {
		results = append(results, &pb.Product{Id: product.Id, Name: product.Name, Price: int32(product.Price)})
	}
	return &pb.ProductsResult{Products: results, Error: nil}, nil
}

/*
コンストラクタ
*/
func NewProductService(rep repository.ProductRepository) *ProductService {
	return &ProductService{rep: rep}
}

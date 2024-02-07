package server

import (
	"context"
	"queryservice/domain/models/products"
	"queryservice/presen/builder"

	"github.com/fullness-MFurukawa/samplepb/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type productServer struct {
	repository products.ProductRepository
	builder    builder.ResultBuilder
	// pb.UnimplementedProductQueryServerをエンベデッドする
	pb.UnimplementedProductQueryServer
}

// コンストラクタ
func NewproductServerImpl(repository products.ProductRepository, builder builder.ResultBuilder) pb.ProductQueryServer {
	return &productServer{repository: repository, builder: builder}
}

// すべての商品を取得して返す(Server streaming RPC)
func (ins *productServer) ListStream(param *emptypb.Empty, stream pb.ProductQuery_ListStreamServer) error {
	if results, err := ins.repository.List(context.Background()); err != nil {
		return err
	} else {
		products := ins.builder.BuildProductsResult(results)
		for _, product := range products.Products {
			// 問合せ結果を送信する
			if err := stream.Send(product); err != nil {
				return err
			}
		}
	}
	return nil
}

// すべての商品を取得して返す
func (ins *productServer) List(ctx context.Context, param *emptypb.Empty) (*pb.ProductsResult, error) {
	if products, err := ins.repository.List(ctx); err != nil {
		return ins.builder.BuildProductsResult(err), nil
	} else {
		return ins.builder.BuildProductsResult(products), nil
	}
}

// 指定されたIDの商品を取得して返す
func (ins *productServer) ById(ctx context.Context, param *pb.ProductParam) (*pb.ProductResult, error) {
	if product, err := ins.repository.FindByProductId(ctx, param.GetId()); err != nil {
		return ins.builder.BuildProductResult(err), nil
	} else {
		return ins.builder.BuildProductResult(product), nil
	}
}

// 指定されたキーワードの商品を取得して返す
func (ins *productServer) ByKeyword(ctx context.Context, param *pb.ProductParam) (*pb.ProductsResult, error) {
	if products, err := ins.repository.FindByProductNameLike(ctx, param.GetKeyword()); err != nil {
		return ins.builder.BuildProductsResult(err), nil
	} else {
		return ins.builder.BuildProductsResult(products), nil
	}
}

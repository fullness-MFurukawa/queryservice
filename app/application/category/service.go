package category

import (
	"context"
	"fmt"
	"queryservice/apperror"
	"queryservice/application/pb"
	"queryservice/domain/repository"
)

// *****
// 　商品カテゴリサービス
// *****
type CategoryService struct {
	rep repository.CategoryRepository
	// 生成したUnimplementedCategtoryServiceServerを埋め込む
	pb.UnimplementedCategtoryServiceServer
}

/*
すべてのカテゴリを取得して返す
*/
func (svr *CategoryService) List(ctx context.Context, param *pb.CategoryParam) (*pb.CategoriesResult, error) {
	categories, err := svr.rep.FindAll()
	if err != nil {
		return &pb.CategoriesResult{Error: &pb.Error{Message: err.Error()}, Categories: nil}, nil
	}
	results := make([]*pb.Category, 0, len(categories))
	for _, category := range categories {
		results = append(results, &pb.Category{Id: category.Id, Name: category.Name})
	}
	return &pb.CategoriesResult{Error: nil, Categories: results}, nil
}

/*
指定されたIDのカテゴリを取得して返す
*/
func (svr *CategoryService) ById(ctx context.Context, param *pb.CategoryParam) (*pb.CategoryResult, error) {
	category, err := svr.rep.FindById(param.GetId())
	if err != nil {
		return &pb.CategoryResult{Result: &pb.CategoryResult_Error{Error: &pb.Error{Message: err.Error()}}}, nil
	}
	//　指定された商品カテゴリIdに該当するカテゴリが見つからない
	if category.Id == "" && category.Name == "" {
		// エラーを生成する
		err := apperror.NewAppError("404", fmt.Sprintf("id[%s]に該当する商品カテゴリが見つかりませんでした。", param.GetId()))
		return &pb.CategoryResult{Result: &pb.CategoryResult_Error{Error: &pb.Error{Message: err.Error()}}}, nil
	}
	result := pb.Category{Id: category.Id, Name: category.Name}
	return &pb.CategoryResult{Result: &pb.CategoryResult_Category{Category: &result}}, nil
}

/*
コンストラクタ
*/
func NewCategoryService(rep repository.CategoryRepository) *CategoryService {
	return &CategoryService{rep: rep}
}

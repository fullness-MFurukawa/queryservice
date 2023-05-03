package exe

import (
	"context"
	"fmt"
	"queryservice/application/pb"
)

type CategoryExecute struct {
	client pb.CategoryServiceClient
}

func (exe *CategoryExecute) List() {
	param := &pb.CategoryParam{} // パラメータの生成
	resp, err := exe.client.List(context.Background(), param)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	for _, category := range resp.GetCategories() {
		fmt.Println(category)
	}
}

/*
コンストラクタ
*/
func NewCategoryExecute(client pb.CategoryServiceClient) *CategoryExecute {
	return &CategoryExecute{client: client}
}

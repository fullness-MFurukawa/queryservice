package exe

import (
	"context"
	"fmt"
	"queryservice/application/pb"
)

type ProductExecute struct {
	client pb.ProductServiceClient
}

func (exe *ProductExecute) List() {
	param := &pb.ProductParam{}
	resp, err := exe.client.List(context.Background(), param)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	for _, product := range resp.GetProducts() {
		fmt.Println(product)
	}
}

/*
コンストラクタ
*/
func NewProductExecute(client pb.ProductServiceClient) *ProductExecute {
	return &ProductExecute{client: client}
}

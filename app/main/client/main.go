package main

import (
	"queryservice/application/pb"
	"queryservice/main/client/exe"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
gRPCクライアント
*/
func main() {
	// gRPCサーバーとのコネクションを確立する
	server := "localhost:8083"
	conn, err := grpc.Dial(
		server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	category := exe.NewCategoryExecute(pb.NewCategoryServiceClient(conn))
	category.List()
	product := exe.NewProductExecute(pb.NewProductServiceClient(conn))
	product.List()
}

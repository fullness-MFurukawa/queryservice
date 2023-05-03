package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"queryservice/application/provider"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	srv "queryservice/application/pb"
)

/*
gRPCサーバーの生成と起動
*/
func main() {

	// アプリケーションサービスを生成する
	service := provider.NewServiceProvider()

	// ポート8083のLisnterを生成する
	port := 8083
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	// gRPCサーバを生成する
	server := grpc.NewServer()

	// サーバーにアプリケーションサービスを登録する
	srv.RegisterCategoryServiceServer(server, &service.CategoryService)
	srv.RegisterProductServiceServer(server, &service.ProductService)

	// gRPCサーバのリフレクションの設定
	reflection.Register(server)

	// gRPCサーバーを起動する
	go func() {
		log.Printf("QueryServiceを起動しました。 ポート:%v", port)
		server.Serve(listener)
	}()

	// Ctrl+Cが入力されたらgRPCサーバを停止する
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("QueryServiceを停止しました。")
	server.GracefulStop()
}

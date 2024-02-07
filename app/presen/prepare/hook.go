package prepare

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc/reflection"
)

func QueryServiceLifecycle(lifecycle fx.Lifecycle, server *QueryServer) {
	lifecycle.Append(fx.Hook{
		// fx起動時の処理
		OnStart: func(ctx context.Context) error {
			// 8083ポートを利用するListenerを生成する
			port := 8083
			listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
			if err != nil {
				return err
			}
			// サーバリフレクションの設定
			reflection.Register(server.Server)
			// 作成したgRPCサーバを起動する
			go func() {
				log.Printf("Query Server 開始 ポート番号: %v", port)
				server.Server.Serve(listener)
			}()
			return nil
		},
		// fx終了時の処理
		OnStop: func(ctx context.Context) error {
			server.Server.GracefulStop() // gRPCサーバを停止する
			log.Println("Query Server 停止")
			return nil
		},
	})
}

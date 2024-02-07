package interceptor

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// ログ出力インターセプタ
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// ログメッセージの出力
	log.Printf("インターセプタログ:%s:%s", info.Server, info.FullMethod)
	// RPC呼び出しを実行
	if resp, err := handler(ctx, req); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

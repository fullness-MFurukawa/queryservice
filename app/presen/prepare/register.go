package prepare

import (
	"context"
	"crypto/tls"
	"embed"
	"queryservice/presen/interceptor"

	"github.com/fullness-MFurukawa/samplepb/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// 埋め込むファイル指定
//
//go:embed queryservice.pem queryservice-key.pem
var embeddedFiles embed.FS

// gRPCサーバの生成とQueryServiceの登録
type QueryServer struct {
	Server *grpc.Server // gRPCServer
}

// インターセプタをチェーン化して実行する
func chainUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// LoggingInterceptor内部でUUIDValidationInterceptorと
	// handlerを呼び出す新しいhandlerを作成する
	newHandler := func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
		// UUID形式の検証インターセプタを実行する
		return interceptor.UUIDValidationInterceptor(currentCtx, currentReq, info, handler)
	}
	// ログ出力インターセプタを実行する
	return interceptor.LoggingInterceptor(ctx, req, info, newHandler)
}

// 認証情報を生成する
func createCreds() credentials.TransportCredentials {
	cert, err := embeddedFiles.ReadFile("queryservice.pem")
	if err != nil {
		panic(err)
	}
	key, err := embeddedFiles.ReadFile("queryservice-key.pem")
	if err != nil {
		panic(err)
	}
	// 証明書と秘密鍵をロードする
	certificate, err := tls.X509KeyPair(cert, key)
	if err != nil {
		panic(err)
	}
	// 資格情報を生成する
	creds := credentials.NewServerTLSFromCert(&certificate)
	return creds
}

// コンストラクタ
func NewQueryServer(category pb.CategoryQueryServer, product pb.ProductQueryServer) *QueryServer {
	serverOpts := []grpc.ServerOption{
		// セキュア伝送サポート機能を登録
		grpc.Creds(createCreds()),
		// ログ出力、入力値検証インターセプタを登録
		grpc.UnaryInterceptor(chainUnaryInterceptor),
	}
	// gRPCサーバを生成する(インターセプタの追加)
	server := grpc.NewServer(serverOpts...)

	// CategoryQueryServerを登録する
	pb.RegisterCategoryQueryServer(server, category)
	// ProductQueryServerを登録する
	pb.RegisterProductQueryServer(server, product)
	return &QueryServer{Server: server}
}

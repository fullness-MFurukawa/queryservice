package presen

import (
	"queryservice/infra/gorm"
	"queryservice/presen/builder"
	"queryservice/presen/prepare"
	"queryservice/presen/server"

	"go.uber.org/fx"
)

var QueryDepend = fx.Options(
	gorm.RepDepend,
	// プレゼンテーション層の依存定義
	fx.Provide(
		builder.NewresultBuilderImpl,
		server.NewcategoryServer,
		server.NewproductServerImpl,
		prepare.NewQueryServer,
	),
	// メソッドの起動
	fx.Invoke(prepare.QueryServiceLifecycle),
)

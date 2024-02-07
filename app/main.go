package main

import (
	"queryservice/presen"

	"go.uber.org/fx"
)

func main() {
	// fxを起動する
	fx.New(
		presen.QueryDepend, // 依存性を定義する
	).Run()
}

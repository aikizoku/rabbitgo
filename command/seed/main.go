package main

import (
	"context"

	"github.com/aikizoku/rabbitgo/command/seed/content"
	"github.dena.jp/esports/toname-backend/command/common"
)

func main() {
	ctx := context.Background()

	// env.jsonの読み込み
	env := common.LoadEnvFile()

	// Inject
	fCli := common.NewFirestoreClient(env.Credentials.Staging)

	u := &content.Sample{
		FCli: fCli,
	}

	// 実行
	u.Generate(ctx)
}

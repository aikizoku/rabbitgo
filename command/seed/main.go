package main

import (
	"context"

	"github.com/aikizoku/rabbitgo/command/lib"
	"github.com/aikizoku/rabbitgo/command/seed/content"
	"github.com/rabee-inc/go-pkg/cloudfirestore"
)

func main() {
	ctx := context.Background()

	// env.jsonの読み込み
	pID := lib.GetProjectID(lib.Staging)

	// Inject
	fCli := cloudfirestore.NewClient(pID)

	u := &content.Sample{
		FCli: fCli,
	}

	// 実行
	u.Generate(ctx)
}

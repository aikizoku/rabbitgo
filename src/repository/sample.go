package repository

import "context"

// Sample ... リポジトリのインターフェースを定義
type Sample interface {
	Hoge(ctx context.Context)
	TestPut(ctx context.Context)
	TestGet(ctx context.Context)
	TestDelete(ctx context.Context)
}

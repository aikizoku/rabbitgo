package repository

import "context"

// Sample ... リポジトリのインターフェースを定義
type Sample interface {
	Hoge(ctx context.Context)
}

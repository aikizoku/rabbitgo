package service

import "context"

// Sample ... サービスのインターフェースを定義
type Sample interface {
	Hoge(ctx context.Context)
}

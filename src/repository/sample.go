package repository

import (
	"context"
)

// Sample ... リポジトリのインターフェース
type Sample interface {
	Sample(ctx context.Context) error
}

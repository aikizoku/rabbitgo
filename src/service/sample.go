package service

import (
	"context"
)

// Sample ... サービスのインターフェース
type Sample interface {
	Sample(ctx context.Context) error
}

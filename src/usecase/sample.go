package usecase

import "context"

// Sample ... ユースケースのインターフェイス
type Sample interface {
	Sample(ctx context.Context) error
}

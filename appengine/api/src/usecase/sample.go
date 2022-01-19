package usecase

import (
	"context"
)

type Sample interface {
	Sample(ctx context.Context) error
}

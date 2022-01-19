package service

import (
	"context"
)

type Sample interface {
	Sample(ctx context.Context) error
}

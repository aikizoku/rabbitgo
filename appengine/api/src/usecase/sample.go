package usecase

import (
	"context"
)

type Sample interface {
	Sample(
		ctx context.Context,
	) error

	UnitTestMethod(
		ctx context.Context,
		hoge int,
		fuga int,
	) (int, error)
}

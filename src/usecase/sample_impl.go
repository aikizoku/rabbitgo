package usecase

import (
	"context"

	"github.com/abyssparanoia/merlin/src/repository"
)

type sample struct {
	repo repository.Sample
}

func (u *sample) Sample(ctx context.Context) error {
	return nil
}

// NewSample ... ユースケースを作成する
func NewSample(repo repository.Sample) Sample {
	return &sample{
		repo: repo,
	}
}

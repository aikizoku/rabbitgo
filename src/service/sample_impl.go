package service

import (
	"context"

	"github.com/aikizoku/merlin/src/repository"
	"github.com/aikizoku/merlin/src/usecase"
)

type sample struct {
	use  usecase.Sample
	repo repository.Sample
}

func (s *sample) Sample(ctx context.Context) error {
	return nil
}

// NewSample ... サービスを作成する
func NewSample(use usecase.Sample, repo repository.Sample) Sample {
	return &sample{
		use:  use,
		repo: repo,
	}
}

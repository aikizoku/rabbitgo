package service

import (
	"context"

	"github.com/aikizoku/skgo/src/repository"
)

type sample struct {
	repo repository.Sample
}

func (s *sample) Sample(ctx context.Context) error {
	return nil
}

// NewSample ... サンプルサービスを取得する
func NewSample(repo repository.Sample) Sample {
	return &sample{
		repo: repo,
	}
}

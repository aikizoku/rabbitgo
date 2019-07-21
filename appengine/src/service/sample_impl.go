package service

import (
	"context"

	"github.com/aikizoku/rabbitgo/appengine/src/repository"
)

type sample struct {
	repo repository.Sample
}

func (s *sample) Sample(ctx context.Context) error {
	return nil
}

// NewSample ... サービスを作成する
func NewSample(repo repository.Sample) Sample {
	return &sample{
		repo: repo,
	}
}

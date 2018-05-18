package service

import (
	"context"

	"github.com/aikizoku/go-gae-template/src/repository"
)

type sample struct {
	repo repository.Sample
}

func NewSample(repo repository.Sample) Sample {
	return &sample{
		repo: repo,
	}
}

func (s *sample) Hoge(ctx context.Context) {

}

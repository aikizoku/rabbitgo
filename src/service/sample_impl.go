package service

import (
	"context"

	"github.com/aikizoku/go-gae-template/src/repository"
	"google.golang.org/appengine/log"
)

type sample struct {
	repo repository.Sample
}

func (s *sample) Hoge(ctx context.Context) {
	log.Debugf(ctx, "call service hoge")
	s.repo.Hoge(ctx)
}

// NewSample ...
func NewSample(repo repository.Sample) Sample {
	return &sample{
		repo: repo,
	}
}

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

func (s *sample) TestPut(ctx context.Context) {
	s.repo.TestPut(ctx)
}

func (s *sample) TestGet(ctx context.Context) {
	s.repo.TestGet(ctx)
}

func (s *sample) TestDelete(ctx context.Context) {
	s.repo.TestDelete(ctx)
}

// NewSample ...
func NewSample(repo repository.Sample) Sample {
	return &sample{
		repo: repo,
	}
}

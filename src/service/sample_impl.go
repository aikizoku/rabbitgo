package service

import (
	"context"

	"github.com/aikizoku/go-gae-template/src/repository"
	"google.golang.org/appengine/log"
)

type sample struct {
	Repository repository.Sample
}

func NewSample(repo repository.Sample) Sample {
	return &sample{
		Repository: repo,
	}
}

func (s *sample) Hoge(ctx context.Context) {
	log.Debugf(ctx, "call service hoge")
	s.Repository.Hoge(ctx)
}

package service

import (
	"context"

	"github.com/rabee-inc/go-pkg/log"

	"github.com/aikizoku/rabbitgo/appengine/api/src/repository"
)

type sample struct {
	rSample repository.Sample
}

func (s *sample) Sample(ctx context.Context) error {
	err := s.rSample.Sample(ctx)
	if err != nil {
		log.Error(ctx, err)
		return err
	}
	return nil
}

// NewSample ... サービスを作成する
func NewSample(rSample repository.Sample) Sample {
	return &sample{
		rSample: rSample,
	}
}

package service

import (
	"context"

	"github.com/rabee-inc/go-pkg/log"

	"github.com/aikizoku/rabbitgo/appengine/api/src/repository"
	"github.com/aikizoku/rabbitgo/appengine/api/src/usecase"
)

type sample struct {
	uSample usecase.Sample
	rSample repository.Sample
}

func NewSample(uSample usecase.Sample, rSample repository.Sample) Sample {
	return &sample{
		uSample,
		rSample,
	}
}

func (s *sample) Sample(ctx context.Context) error {
	err := s.rSample.Sample(ctx)
	if err != nil {
		log.Error(ctx, err)
		return err
	}
	return nil
}

package usecase

import (
	"context"
	"fmt"

	"github.com/rabee-inc/go-pkg/log"

	"github.com/aikizoku/rabbitgo/appengine/api/src/repository"
)

type sample struct {
	rSample repository.Sample
}

func NewSample(rSample repository.Sample) Sample {
	return &sample{
		rSample: rSample,
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

func (s *sample) UnitTestMethod(
	ctx context.Context,
	hoge int,
	fuga int,
) (int, error) {
	if hoge < 0 || fuga < 0 {
		return 0, fmt.Errorf("hoge or fuga is less than 0")
	}
	return hoge + fuga, nil
}

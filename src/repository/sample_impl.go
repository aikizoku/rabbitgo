package repository

import (
	"context"
)

type sample struct {
}

func NewSample() Sample {
	return &sample{}
}

func (s *sample) Hoge(ctx context.Context) {
}

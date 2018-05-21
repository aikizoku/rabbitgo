package repository

import (
	"context"

	"google.golang.org/appengine/log"
)

type sample struct {
}

func NewSample() Sample {
	return &sample{}
}

func (s *sample) Hoge(ctx context.Context) {
	log.Debugf(ctx, "call repository hoge")
}

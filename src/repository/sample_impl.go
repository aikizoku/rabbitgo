package repository

import (
	"context"

	"github.com/aikizoku/go-gae-template/src/infrastructure"
	"google.golang.org/appengine/log"
)

type sample struct {
	Http infrastructure.HTTP
}

func NewSample(http infrastructure.HTTP) Sample {
	return &sample{
		Http: http,
	}
}

func (s *sample) Hoge(ctx context.Context) {
	log.Debugf(ctx, "call repository hoge")
}

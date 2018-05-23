package repository

import (
	"context"

	"github.com/aikizoku/go-gae-template/src/infrastructure"
	"google.golang.org/appengine/log"
)

type sample struct {
	http infrastructure.HTTP
}

func (s *sample) Hoge(ctx context.Context) {
	log.Debugf(ctx, "call repository hoge")
}

// NewSample ...
func NewSample(http infrastructure.HTTP) Sample {
	return &sample{
		http: http,
	}
}

package repository

import (
	"context"
)

type sample struct {
}

func (r *sample) Sample(ctx context.Context) error {
	return nil
}

func (r *sample) HTTPPost(ctx context.Context) error {
	return nil
}

// NewSample ... リポジトリを作成する
func NewSample() Sample {
	return &sample{}
}

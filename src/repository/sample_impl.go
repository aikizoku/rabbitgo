package repository

import (
	"context"

	"cloud.google.com/go/firestore"
)

type sample struct {
	client *firestore.Client
}

func (r *sample) Sample(ctx context.Context) error {
	return nil
}

func (r *sample) HTTPPost(ctx context.Context) error {
	return nil
}

// NewSample ... サンプルリポジトリを取得する
func NewSample(client *firestore.Client) Sample {
	return &sample{
		client: client,
	}
}

package repository

import (
	"context"

	"cloud.google.com/go/firestore"
)

type sample struct {
	fCli *firestore.Client
}

func (r *sample) Sample(ctx context.Context) error {
	return nil
}

// NewSample ... リポジトリを作成する
func NewSample(fCli *firestore.Client) Sample {
	return &sample{
		fCli: fCli,
	}
}

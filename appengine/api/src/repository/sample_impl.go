package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/rabee-inc/go-pkg/images"
)

type sample struct {
	cFirestore *firestore.Client
	cImages    *images.Client
}

func (r *sample) Sample(ctx context.Context) error {
	return nil
}

// NewSample ... リポジトリを作成する
func NewSample(cFirestore *firestore.Client, cImages *images.Client) Sample {
	return &sample{
		cFirestore: cFirestore,
		cImages:    cImages,
	}
}

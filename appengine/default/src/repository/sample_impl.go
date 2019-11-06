package repository

import (
	"context"

	"cloud.google.com/go/firestore"

	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/images"
)

type sample struct {
	fCli   *firestore.Client
	imgCli *images.Client
}

func (r *sample) Sample(ctx context.Context) error {
	return nil
}

// NewSample ... リポジトリを作成する
func NewSample(fCli *firestore.Client, imgCli *images.Client) Sample {
	return &sample{
		fCli:   fCli,
		imgCli: imgCli,
	}
}

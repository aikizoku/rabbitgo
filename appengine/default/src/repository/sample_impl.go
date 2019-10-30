package repository

import (
	"context"

	"cloud.google.com/go/firestore"

	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/cloudfirestore"
	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/images"
	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/log"
)

type sample struct {
	fCli   *firestore.Client
	imgCli *images.Client
}

func (r *sample) Sample(ctx context.Context) error {
	// ref := r.fCli.Collection("main").Doc("lvr3kT96RUEtw2DsagTI").Collection("sub").Doc("WUGepFY6syZBkW8FKAAT")
	err := r.imgCli.SendConvertRequest(
		ctx,
		"hogehoge",
		"http://shiroichi.sakura.ne.jp/wp-content/uploads/2011/04/largeimage.jpg",
		"image/sample",
		[]*cloudfirestore.DocRef{
			&cloudfirestore.DocRef{CollectionName: "main", DocID: "lvr3kT96RUEtw2DsagTI"},
			&cloudfirestore.DocRef{CollectionName: "sub", DocID: "WUGepFY6syZBkW8FKAAT"},
		},
		"profile_image")
	if err != nil {
		log.Errorm(ctx, "r.imgCli.SendConvertRequest", err)
		return err
	}
	return nil
}

// NewSample ... リポジトリを作成する
func NewSample(fCli *firestore.Client, imgCli *images.Client) Sample {
	return &sample{
		fCli:   fCli,
		imgCli: imgCli,
	}
}

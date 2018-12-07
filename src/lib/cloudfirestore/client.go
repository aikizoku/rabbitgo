package cloudfirestore

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/aikizoku/beego/src/lib/log"
)

// NewClient ... Firestoreのクライアントを取得する
func NewClient(ctx context.Context) (*firestore.Client, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Errorm(ctx, "firebase.NewApp", err)
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Errorm(ctx, "app.Firestore", err)
		return nil, err
	}
	return client, nil
}

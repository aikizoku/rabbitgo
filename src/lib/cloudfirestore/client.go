package cloudfirestore

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/aikizoku/skgo/src/lib/log"
	"google.golang.org/api/option"
)

// NewClient ... Firestoreのクライアントを取得する
func NewClient(credentialsPath string) (*firestore.Client, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(ctx, nil, opt)
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

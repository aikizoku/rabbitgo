package cloudfirestore

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// NewClient ... Firestoreのクライアントを取得する
func NewClient(credentialsPath string) *firestore.Client {
	ctx := context.Background()
	opt := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}
	cli, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}
	return cli
}

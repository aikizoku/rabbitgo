package cloudfirestore

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// NewClient ... Firestoreのクライアントを取得する
func NewClient(credentialsPath string) *firestore.Client {
	ctx := context.Background()
	cOpt := option.WithCredentialsFile(credentialsPath)
	gOpt := option.WithGRPCDialOption(grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                30 * time.Millisecond,
		Timeout:             20 * time.Millisecond,
		PermitWithoutStream: true,
	}))
	app, err := firebase.NewApp(ctx, nil, cOpt, gOpt)
	if err != nil {
		panic(err)
	}
	cli, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}
	return cli
}

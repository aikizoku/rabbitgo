package firebaseauth

import (
	"context"
	"fmt"

	"firebase.google.com/go/auth"

	"github.com/aikizoku/rabbitgo/appengine/src/lib/log"
)

type service struct {
	cli *auth.Client
}

// Authentication ... 認証を行う
func (s *service) Authentication(ctx context.Context, ah string) (string, *Claims, error) {
	var userID string
	claims := &Claims{}

	token := getTokenByAuthHeader(ah)
	if token == "" {
		err := log.Warninge(ctx, "token empty error")
		return userID, claims, err
	}

	t, err := s.cli.VerifyIDToken(ctx, token)
	if err != nil {
		msg := fmt.Sprintf("c.VerifyIDToken: %s", token)
		log.Warningm(ctx, msg, err)
		return userID, claims, err
	}

	userID = t.UID
	claims.SetMap(t.Claims)

	return userID, claims, nil
}

// SetCustomClaims ... カスタムClaimsを設定
func (s *service) SetCustomClaims(ctx context.Context, userID string, claims *Claims) error {
	err := s.cli.SetCustomUserClaims(ctx, userID, claims.ToMap())
	if err != nil {
		log.Errorm(ctx, "c.SetCustomUserClaims", err)
		return err
	}
	return nil
}

// NewService ... Serviceを作成する
func NewService(cli *auth.Client) Service {
	return &service{
		cli: cli,
	}
}

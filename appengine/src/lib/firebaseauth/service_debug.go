package firebaseauth

import (
	"context"
	"fmt"

	"firebase.google.com/go/auth"

	"github.com/aikizoku/rabbitgo/appengine/src/lib/log"
)

type serviceDebug struct {
	cli *auth.Client
}

// Authentication ... 認証を行う
func (s *serviceDebug) Authentication(ctx context.Context, ah string) (string, *Claims, error) {
	var userID string
	claims := &Claims{}

	// ユーザーを取得できたらデバッグリクエストと判定する
	if user := getUserByAuthHeader(ah); user != "" {
		claims = NewDummyClaims()
		return user, claims, nil
	}

	// 通常の認証を行う
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
func (s *serviceDebug) SetCustomClaims(ctx context.Context, userID string, claims *Claims) error {
	// AuthorizationHeaderからUserが取得できたらデバッグリクエストと判定する
	ah := getAuthHeader(ctx)
	if getUserByAuthHeader(ah) == "" {
		err := s.cli.SetCustomUserClaims(ctx, userID, claims.ToMap())
		if err != nil {
			log.Errorm(ctx, "c.SetCustomUserClaims", err)
			return err
		}
	}
	return nil
}

// NewDebugService ... DebugServiceを作成する
func NewDebugService(cli *auth.Client) Service {
	return &serviceDebug{
		cli: cli,
	}
}

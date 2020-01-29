package firebaseauth

import (
	"context"
	"fmt"

	"firebase.google.com/go/auth"

	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/log"
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

func (s *serviceDebug) GetEmail(ctx context.Context, userID string) (string, error) {
	// ユーザーを取得できたらデバッグリクエストと判定する
	ah := getAuthHeader(ctx)
	if user := getUserByAuthHeader(ah); user != "" {
		return "hirose.yuuki@rabee.jp", nil
	}

	// FirebaseAuthUserを取得
	user, err := s.cli.GetUser(ctx, userID)
	if err != nil {
		log.Errorm(ctx, "s.cli.GetUser", err)
		return "", err
	}
	if user == nil {
		return "", err
	}
	return user.Email, nil
}

func (s *serviceDebug) GetTwitterID(ctx context.Context, userID string) (string, error) {
	// AuthorizationHeaderからUserが取得できたらデバッグリクエストと判定する
	ah := getAuthHeader(ctx)
	if getUserByAuthHeader(ah) != "" {
		return "", nil
	}

	// FirebaseAuthUserを取得
	user, err := s.cli.GetUser(ctx, userID)
	if err != nil {
		log.Errorm(ctx, "s.cli.GetUser", err)
		return "", err
	}
	if user == nil {
		return "", err
	}

	dst := ""
	for _, userInfo := range user.ProviderUserInfo {
		if userInfo != nil && userInfo.ProviderID == "twitter.com" {
			dst = userInfo.UID
			break
		}
	}
	return dst, nil
}

// NewDebugService ... DebugServiceを作成する
func NewDebugService(cli *auth.Client) Service {
	return &serviceDebug{
		cli: cli,
	}
}

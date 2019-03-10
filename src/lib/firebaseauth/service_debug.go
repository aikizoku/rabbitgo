package firebaseauth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/aikizoku/skgo/src/lib/log"
)

const (
	debugTokenPrefix string = "user="
)

type serviceDebug struct {
}

// SetCustomClaims ... カスタムClaimsを設定
func (s *serviceDebug) SetCustomClaims(ctx context.Context, userID string, claims Claims) error {
	c, err := s.getAuthClient(ctx)
	if err != nil {
		log.Errorm(ctx, "s.getAuthClient", err)
		return err
	}

	err = c.SetCustomUserClaims(ctx, userID, claims.ToMap())
	if err != nil {
		log.Errorm(ctx, "c.SetCustomUserClaims", err)
		return err
	}

	return nil
}

// Authentication ... 認証を行う
func (s *serviceDebug) Authentication(ctx context.Context, r *http.Request) (string, Claims, error) {
	var userID string
	claims := Claims{}

	token := s.getTokenByRequest(r)
	if token == "" {
		err := log.Warninge(ctx, "token empty error")
		return userID, claims, err
	}

	if strings.HasPrefix(token, debugTokenPrefix) {
		userID = token[len(debugTokenPrefix):]

		/* ダミーで割り当てるカスタムクレームを設定 */
	} else {
		c, err := s.getAuthClient(ctx)
		if err != nil {
			log.Warningm(ctx, "s.getAuthClient", err)
			return userID, claims, err
		}

		t, err := c.VerifyIDToken(ctx, token)
		if err != nil {
			msg := fmt.Sprintf("c.VerifyIDToken: %s", token)
			log.Warningm(ctx, msg, err)
			return userID, claims, err
		}

		userID = t.UID
		claims.SetMap(t.Claims)
	}

	return userID, claims, nil
}

func (s *serviceDebug) getAuthClient(ctx context.Context) (*auth.Client, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Errorm(ctx, "firebase.NewApp", err)
		return nil, err
	}
	c, err := app.Auth(ctx)
	if err != nil {
		log.Errorm(ctx, "app.Auth", err)
		return nil, err
	}
	return c, nil
}

func (s *serviceDebug) getTokenByRequest(r *http.Request) string {
	if ah := r.Header.Get("Authorization"); ah != "" {
		pLen := len(headerPrefix)
		if len(ah) > pLen && strings.ToUpper(ah[0:pLen]) == headerPrefix {
			return ah[pLen+1:]
		}
	}
	return ""
}

// NewDebugService ... DebugServiceを作成する
func NewDebugService() Service {
	return &serviceDebug{}
}

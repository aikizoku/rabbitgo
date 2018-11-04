package firebaseauth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/appengine/log"
)

type service struct {
}

// SetCustomClaims ... カスタムClaimsを設定
func (s *service) SetCustomClaims(ctx context.Context, userID string, claims Claims) error {
	c, err := s.getAuthClient(ctx)
	if err != nil {
		log.Errorf(ctx, "faild to get auth client")
		return err
	}

	err = c.SetCustomUserClaims(ctx, userID, claims.ToMap())
	if err != nil {
		log.Errorf(ctx, err.Error())
		return err
	}

	return nil
}

// Authentication ... 認証を行う
func (s *service) Authentication(ctx context.Context, r *http.Request) (string, Claims, error) {
	var userID string
	claims := Claims{}

	c, err := s.getAuthClient(ctx)
	if err != nil {
		log.Warningf(ctx, "faild to get auth client")
		return userID, claims, err
	}

	idToken := s.getAuthorizationHeader(r)
	if idToken == "" {
		err = fmt.Errorf("no auth token error")
		return userID, claims, err
	}

	t, err := c.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Warningf(ctx, "c.VerifyIDToken: %s", err.Error())
		return userID, claims, err
	}

	userID = t.UID
	claims.SetMap(t.Claims)

	return userID, claims, nil
}

func (s *service) getAuthClient(ctx context.Context) (*auth.Client, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Warningf(ctx, "create firebase app error: %s", err.Error())
		return nil, err
	}
	c, err := app.Auth(ctx)
	if err != nil {
		log.Warningf(ctx, "create auth client error: %s", err.Error())
		return nil, err
	}
	return c, nil
}

func (s *service) getAuthorizationHeader(r *http.Request) string {
	if ah := r.Header.Get("Authorization"); ah != "" {
		if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
			return ah[7:]
		}
	}
	return ""
}

// NewService ... Serviceを作成する
func NewService() Service {
	return &service{}
}

package service

import (
	"context"
	"net/http"

	"github.com/aikizoku/beego/src/model"
)

// Authenticator ... 認証
type Authenticator interface {
	Authentication(ctx context.Context, r *http.Request) (string, model.Claims, error)
	SetCustomClaims(ctx context.Context, userID string, claims model.Claims) error
}

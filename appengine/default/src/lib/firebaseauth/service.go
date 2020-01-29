package firebaseauth

import (
	"context"
)

// Service ... Firebase認証の機能を提供する
type Service interface {
	Authentication(ctx context.Context, ah string) (string, *Claims, error)
	SetCustomClaims(ctx context.Context, userID string, claims *Claims) error
	GetEmail(ctx context.Context, userID string) (string, error)
	GetTwitterID(ctx context.Context, userID string) (string, error)
}

package firebaseauth

import "context"

type contextKey string

const (
	authHeaderContextKey contextKey = "firebaseauth:auth_header"
	userIDContextKey     contextKey = "firebaseauth:user_id"
	claimsContextKey     contextKey = "firebaseauth:claims"
)

func getAuthHeader(ctx context.Context) string {
	return ctx.Value(authHeaderContextKey).(string)
}

// GetUserID ... FirebaseAuthのユーザーIDを取得
func GetUserID(ctx context.Context) string {
	if dst := ctx.Value(userIDContextKey); dst != nil {
		return dst.(string)
	}
	return ""
}

// GetClaims ... FirebaseAuthのJWTClaimsの値を取得
func GetClaims(ctx context.Context) (*Claims, bool) {
	if dst := ctx.Value(claimsContextKey); dst != nil {
		return dst.(*Claims), true
	}
	return nil, false
}

func setAuthHeader(ctx context.Context, ah string) context.Context {
	return context.WithValue(ctx, authHeaderContextKey, ah)
}

func setUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDContextKey, userID)
}

func setClaims(ctx context.Context, claims *Claims) context.Context {
	return context.WithValue(ctx, claimsContextKey, claims)
}

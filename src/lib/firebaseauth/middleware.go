package firebaseauth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aikizoku/merlin/src/lib/log"
	"github.com/aikizoku/merlin/src/lib/renderer"
)

// Middleware ... Firebase認証のミドルウェア
type Middleware struct {
	Svc Service
}

// Handle ... Firebase認証をする
func (m *Middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Headerを取得
		ah := r.Header.Get("Authorization")
		if ah == "" {
			m.renderError(ctx, w, http.StatusForbidden, "no Authorization header")
			return
		}
		ctx = setAuthHeader(ctx, ah)
		log.Debugf(ctx, "Authorization Header: %s", ah)

		// 認証
		userID, claims, err := m.Svc.Authentication(ctx, ah)
		if err != nil {
			m.renderError(ctx, w, http.StatusForbidden, err.Error())
			return
		}

		// 認証結果を設定
		ctx = setUserID(ctx, userID)
		log.Debugf(ctx, "UserID: %s", userID)

		ctx = setClaims(ctx, claims)
		log.Debugf(ctx, "Claims: %v", claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) renderError(ctx context.Context, w http.ResponseWriter, status int, msg string) {
	log.Warningf(ctx, msg)
	renderer.Error(w, status, fmt.Sprintf("%d Authorization failed", status))
}

// NewMiddleware ... Middlewareを作成する
func NewMiddleware(svc Service) *Middleware {
	return &Middleware{
		Svc: svc,
	}
}

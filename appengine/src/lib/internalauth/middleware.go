package internalauth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aikizoku/rabbitgo/appengine/src/lib/log"
	"github.com/aikizoku/rabbitgo/appengine/src/lib/renderer"
)

// Middleware ... 内部認証機能を提供するミドルウェア
type Middleware struct {
	Token string
}

// Handle ... ハンドラ
func (m *Middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ah := r.Header.Get("Authorization")
		if ah == "" || ah != m.Token {
			ctx := r.Context()
			m.renderError(ctx, w, http.StatusForbidden, "Internal auth error token: %s", ah)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) renderError(ctx context.Context, w http.ResponseWriter, status int, format string, args ...string) {
	msg := fmt.Sprintf(format, args)
	log.Warningf(ctx, msg)
	renderer.Error(ctx, w, status, msg)
}

// NewMiddleware ... Middlewareを作成する
func NewMiddleware(token string) *Middleware {
	return &Middleware{
		Token: token,
	}
}

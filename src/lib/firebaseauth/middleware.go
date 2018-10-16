package firebaseauth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aikizoku/beego/src/config"
	"github.com/aikizoku/beego/src/model"
	"github.com/unrolled/render"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// UserIDContextKey ... UserIDのContextKey
const UserIDContextKey config.ContextKey = "user_id"

// ClaimsContextKey ... ClaimsのContextKey
const ClaimsContextKey config.ContextKey = "claims"

// Middleware ... JSONRPC2に準拠したミドルウェア
type Middleware struct {
}

// Auth ... Firebase認証をする
func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		userID, claims, err := authentication(ctx, r)
		if err != nil {
			m.renderError(ctx, w, http.StatusForbidden, "authentication: "+err.Error())
			return
		}

		rctx := r.Context()
		rctx = context.WithValue(rctx, UserIDContextKey, userID)
		rctx = context.WithValue(rctx, ClaimsContextKey, claims)

		log.Debugf(ctx, "UserID: %s", userID)
		log.Debugf(ctx, "Claims: %v", claims)

		next.ServeHTTP(w, r.WithContext(rctx))
	})
}

// DummyAuth ... Firebaseダミー認証をする
func (m *Middleware) DummyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := "DUMMY_USER_ID"
		claims := model.Claims{
			// ダミーのClaimsを設定する
			Sample: "sample",
		}
		rctx := r.Context()
		rctx = context.WithValue(rctx, UserIDContextKey, userID)
		rctx = context.WithValue(rctx, ClaimsContextKey, claims)
		next.ServeHTTP(w, r.WithContext(rctx))
	})
}

func (m *Middleware) renderError(ctx context.Context, w http.ResponseWriter, status int, msg string) {
	log.Warningf(ctx, msg)
	render.New().Text(w, status, fmt.Sprintf("%d %s", status, msg))
}

// NewMiddleware ... Middlewareを作成する
func NewMiddleware() *Middleware {
	return &Middleware{}
}

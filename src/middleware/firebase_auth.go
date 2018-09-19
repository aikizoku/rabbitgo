package middleware

import (
	"context"
	"net/http"

	"github.com/aikizoku/beego/src/config"
	"github.com/aikizoku/beego/src/model"
	"github.com/aikizoku/beego/src/service"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// UserIDContextKey ... UserIDのContextKey
const UserIDContextKey config.ContextKey = "user_id"

// ClaimsContextKey ... ClaimsのContextKey
const ClaimsContextKey config.ContextKey = "claims"

// FirebaseAuth ... Firebase認証
type FirebaseAuth struct {
	Authenticator service.Authenticator
}

// Authentication ... 認証をする
func (m *FirebaseAuth) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		userID, claims, err := m.Authenticator.Authentication(ctx, r)
		if err != nil {
			m.renderError(ctx, w, http.StatusForbidden, "Authentication: "+err.Error())
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

// DummyAuthentication ... ダミー認証をする
func (m *FirebaseAuth) DummyAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := "DUMMY_USER_ID"
		claims := model.Claims{}
		rctx := r.Context()
		rctx = context.WithValue(rctx, UserIDContextKey, userID)
		rctx = context.WithValue(rctx, ClaimsContextKey, claims)
		next.ServeHTTP(w, r.WithContext(rctx))
	})
}

func (m *FirebaseAuth) renderError(ctx context.Context, w http.ResponseWriter, status int, msg string) {
	log.Warningf(ctx, msg)
	RenderError(w, status, msg)
}

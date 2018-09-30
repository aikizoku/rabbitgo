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

// FirebaseAuth ... Firebase認証をする
func FirebaseAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		userID, claims, err := Authentication(ctx, r)
		if err != nil {
			renderError(ctx, w, http.StatusForbidden, "Authentication: "+err.Error())
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

// DummyFirebaseAuth ... Firebaseダミー認証をする
func DummyFirebaseAuth(next http.Handler) http.Handler {
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

func renderError(ctx context.Context, w http.ResponseWriter, status int, msg string) {
	log.Warningf(ctx, msg)
	render.New().Text(w, status, fmt.Sprintf("%d %s", status, msg))
}

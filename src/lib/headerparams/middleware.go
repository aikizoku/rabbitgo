package headerparams

import (
	"context"
	"fmt"
	"net/http"

	"github.com/unrolled/render"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// ContextKey ... ContextKeyの型定義
type ContextKey string

// HeaderParamsContextKey ... HeaderParamsのContextKey
const HeaderParamsContextKey ContextKey = "header_params"

// Middleware ... Headerに関する機能を提供する
type Middleware struct {
	Svc Service
}

// Handle ... リクエストヘッダーのパラメータを取得する
func (m *Middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		h, err := m.Svc.Get(ctx, r)
		if err != nil {
			m.renderError(ctx, w, http.StatusBadRequest, "headerparams.Service.Get: "+err.Error())
			return
		}
		rctx := r.Context()
		rctx = context.WithValue(rctx, HeaderParamsContextKey, h)
		next.ServeHTTP(w, r.WithContext(rctx))
	})
}

func (m *Middleware) renderError(ctx context.Context, w http.ResponseWriter, status int, msg string) {
	log.Warningf(ctx, msg)
	render.New().Text(w, status, fmt.Sprintf("%d %s", status, msg))
}

// NewMiddleware ... Middlewareを作成する
func NewMiddleware(svc Service) *Middleware {
	return &Middleware{
		Svc: svc,
	}
}

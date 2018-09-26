package middleware

import (
	"context"
	"net/http"

	"github.com/aikizoku/beego/src/config"
	"github.com/aikizoku/beego/src/model"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"gopkg.in/go-playground/validator.v9"
)

// HeaderParamsContextKey ... HeaderParamsのContextKey
const HeaderParamsContextKey config.ContextKey = "header_params"

const (
	headerKeySample string = "X-Sample"
)

// GetHeaderParams ... リクエストヘッダーのパラメータを取得する
func GetHeaderParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		h := model.HeaderParams{
			Sample: r.Header.Get(headerKeySample),
		}
		v := validator.New()
		if err := v.Struct(h); err != nil {
			log.Warningf(ctx, err.Error())
			RenderError(w, http.StatusBadRequest, err.Error())
			return
		}
		rctx := r.Context()
		rctx = context.WithValue(rctx, HeaderParamsContextKey, h)
		next.ServeHTTP(w, r.WithContext(rctx))
	})
}

// GetDummyHeaderParams ... リクエストヘッダーのダミーパラメータを取得する
func GetDummyHeaderParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		h := model.HeaderParams{
			// ダミーのHeaderParamsを設定する
			Sample: "sample",
		}
		v := validator.New()
		if err := v.Struct(h); err != nil {
			log.Warningf(ctx, err.Error())
			RenderError(w, http.StatusBadRequest, err.Error())
			return
		}
		rctx := r.Context()
		rctx = context.WithValue(rctx, HeaderParamsContextKey, h)
		next.ServeHTTP(w, r.WithContext(rctx))
	})
}

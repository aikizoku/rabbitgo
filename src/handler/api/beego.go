package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/aikizoku/beego/src/middleware"
	"github.com/aikizoku/beego/src/model"
	"github.com/aikizoku/beego/src/service"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// BeegoHandler ... 記事のハンドラ
type BeegoHandler struct {
	Beego service.Beego
}

// Beego ... サンプルハンドラ
func (h *BeegoHandler) Beego(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// HTTPHeaderの値を取得
	headerParam := r.Context().Value(middleware.HeaderParamsContextKey).(model.HeaderParams).Beego
	if headerParam == "" {
		h.handleError(ctx, w, http.StatusBadRequest, "invalid header param is empty")
		return
	}

	// URLParamの値を取得
	urlParam := chi.URLParam(r, "beego")
	if urlParam == "" {
		h.handleError(ctx, w, http.StatusBadRequest, "invalid url param is empty")
		return
	}

	// フォームの値を取得
	formParam := r.FormValue("beego")
	if formParam == "" {
		h.handleError(ctx, w, http.StatusBadRequest, "invalid form param is empty")
		return
	}

	// FirebaseAuthのユーザーIDを取得
	userID := r.Context().Value(middleware.UserIDContextKey).(string)

	// FirebaseAuthのJWTClaimsの値を取得
	claims := r.Context().Value(middleware.ClaimsContextKey).(model.Claims)

	// Serviceを実行する
	beego, err := h.Service.Beego(ctx)
	if err != nil {
		h.handleError(ctx, w, http.StatusInternalServerError, "h.Service.Beego: "+err.Error())
		return
	}

	middleware.RenderJSON(w, http.StatusOK, struct {
		Beego model.Beego `json:"beego"`
		Hoge  string      `json:"hoge,omitempty"`
	}{
		Beego: beego,
		Hoge:  "",
	})
}

func (h *BeegoHandler) handleError(ctx context.Context, w http.ResponseWriter, status int, msg string) {
	log.Errorf(ctx, msg)
	middleware.RenderError(w, status, msg)
}

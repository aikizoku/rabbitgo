package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/aikizoku/beego/src/lib/firebaseauth"
	"github.com/aikizoku/beego/src/middleware"
	"github.com/aikizoku/beego/src/model"
	"github.com/aikizoku/beego/src/service"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// SampleHandler ... 記事のハンドラ
type SampleHandler struct {
	Svc service.Sample
}

// Sample ... サンプルハンドラ
func (h *SampleHandler) Sample(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// HTTPHeaderの値を取得
	headerParams := r.Context().Value(middleware.HeaderParamsContextKey).(model.HeaderParams)
	log.Debugf(ctx, "HeaderParams: %v", headerParams)

	// URLParamの値を取得
	urlParam := chi.URLParam(r, "sample")
	if urlParam == "" {
		h.handleError(ctx, w, http.StatusBadRequest, "invalid url param is empty")
		return
	}
	log.Debugf(ctx, "URLParam: %s", urlParam)

	// フォームの値を取得
	formParam := r.FormValue("sample")
	if formParam == "" {
		h.handleError(ctx, w, http.StatusBadRequest, "invalid form param is empty")
		return
	}
	log.Debugf(ctx, "FormParams: %s", formParam)

	// FirebaseAuthのユーザーIDを取得
	userID := r.Context().Value(firebaseauth.UserIDContextKey).(string)
	log.Debugf(ctx, "UserID: %s", userID)

	// FirebaseAuthのJWTClaimsの値を取得
	claims := r.Context().Value(firebaseauth.ClaimsContextKey).(model.Claims)
	log.Debugf(ctx, "Claims: %v", claims)

	// Serviceを実行する
	sample, err := h.Svc.Sample(ctx)
	if err != nil {
		h.handleError(ctx, w, http.StatusInternalServerError, "h.Service.Sample: "+err.Error())
		return
	}

	middleware.RenderJSON(w, http.StatusOK, struct {
		Sample model.Sample `json:"sample"`
		Hoge   string       `json:"hoge,omitempty"`
	}{
		Sample: sample,
		Hoge:   "",
	})
}

// TestDataStore ... DataStoreテスト
func (h *SampleHandler) TestDataStore(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	err := h.Svc.TestDataStore(ctx)
	if err != nil {
		h.handleError(ctx, w, http.StatusInternalServerError, "h.Svc.TestDataStore: "+err.Error())
		return
	}

	middleware.RenderSuccess(w)
}

// TestCloudSQL ... CloudSQLテスト
func (h *SampleHandler) TestCloudSQL(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	err := h.Svc.TestCloudSQL(ctx)
	if err != nil {
		h.handleError(ctx, w, http.StatusInternalServerError, "h.Svc.TestCloudSQL: "+err.Error())
		return
	}

	middleware.RenderSuccess(w)
}

// TestHTTP ... HTTPテスト
func (h *SampleHandler) TestHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	err := h.Svc.TestHTTP(ctx)
	if err != nil {
		h.handleError(ctx, w, http.StatusInternalServerError, "h.Svc.TestHTTP: "+err.Error())
		return
	}

	middleware.RenderSuccess(w)
}

func (h *SampleHandler) handleError(ctx context.Context, w http.ResponseWriter, status int, msg string) {
	log.Errorf(ctx, msg)
	middleware.RenderError(w, status, msg)
}

package worker

import (
	"net/http"

	"github.com/aikizoku/beego/src/middleware"
	"github.com/aikizoku/beego/src/service"
	"google.golang.org/appengine"
)

// SampleHandler ... サンプルのハンドラ定義
type SampleHandler struct {
	Svc service.SampleService
}

// Beegos ... サンプルのハンドラ
func (h *SampleHandler) Beegos(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	middleware.RenderSuccess(w)
}

// Beego ... サンプルのハンドラ
func (h *SampleHandler) Beego(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	middleware.RenderSuccess(w)
}

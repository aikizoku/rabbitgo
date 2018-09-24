package worker

import (
	"net/http"

	"github.com/aikizoku/beego/src/middleware"
	"github.com/aikizoku/beego/src/service"
	"google.golang.org/appengine"
)

// BeegoHandler ... サンプルのハンドラ定義
type BeegoHandler struct {
	Beego service.Beego
}

// Beegos ... サンプルのハンドラ
func (h *BeegoHandler) Beegos(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	middleware.RenderSuccess(w)
}

// Beego ... サンプルのハンドラ
func (h *BeegoHandler) Beego(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	middleware.RenderSuccess(w)
}



itms-apps://itunes.apple.com/jp/app/id31241231
appID	String	"31241231"	
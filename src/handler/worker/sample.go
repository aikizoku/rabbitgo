package worker

import (
	"net/http"

	"github.com/aikizoku/beego/src/middleware"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// SampleHandler ... サンプルのハンドラ定義
type SampleHandler struct {
}

// CronHandler ... Cronから実行されるハンドラ
func (h *SampleHandler) CronHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "call cron handler")
	middleware.RenderSuccess(w)
}

// TaskQueueHandler ... TaskQueueで実行されるハンドラ
func (h *SampleHandler) TaskQueueHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "call task queue handler")
	middleware.RenderSuccess(w)
}

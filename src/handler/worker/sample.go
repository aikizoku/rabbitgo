package worker

import (
	"net/http"

	"github.com/aikizoku/gocci/src/handler"
	"github.com/aikizoku/gocci/src/lib/log"
)

// SampleHandler ... サンプルのハンドラ定義
type SampleHandler struct {
}

// Cron ... Cronから実行されるハンドラ
func (h *SampleHandler) Cron(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Debugf(ctx, "Cronから実行される")

	handler.RenderSuccess(w)
}

// TaskQueue ... TaskQueueで実行されるハンドラ
func (h *SampleHandler) TaskQueue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Debugf(ctx, "TaskQueueから実行される")

	handler.RenderSuccess(w)
}

// NewSampleHandler ... SampleHandlerを作成する
func NewSampleHandler() *SampleHandler {
	return &SampleHandler{}
}

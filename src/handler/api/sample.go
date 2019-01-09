package api

import (
	"context"
	"net/http"

	"github.com/aikizoku/gocci/src/handler"
	"github.com/aikizoku/gocci/src/lib/log"
	"github.com/aikizoku/gocci/src/service"
)

// SampleHandler ... サンプルのハンドラ
type SampleHandler struct {
	Svc service.Sample
}

// Sample ... サンプルハンドラ
func (h *SampleHandler) Sample(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := h.Svc.Sample(ctx)
	if err != nil {
		h.handleError(ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.RenderSuccess(w)
}

func (h *SampleHandler) handleError(ctx context.Context, w http.ResponseWriter, status int, msg string) {
	log.Errorf(ctx, msg)
	handler.RenderError(w, status, msg)
}

// NewSampleHandler ... SampleHandlerを作成する
func NewSampleHandler(svc service.Sample) *SampleHandler {
	return &SampleHandler{
		Svc: svc,
	}
}

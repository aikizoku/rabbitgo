package api

import (
	"net/http"

	"github.com/aikizoku/merlin/src/handler"
	"github.com/aikizoku/merlin/src/service"
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
		handler.HandleError(ctx, w, "h.Svc.Sample", err)
		return
	}

	handler.RenderSuccess(w)
}

// NewSampleHandler ... SampleHandlerを作成する
func NewSampleHandler(svc service.Sample) *SampleHandler {
	return &SampleHandler{
		Svc: svc,
	}
}

package site

import (
	"net/http"

	"github.com/rabee-inc/go-pkg/renderer"

	"github.com/aikizoku/rabbitgo/appengine/api/src/service"
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
		renderer.HandleError(ctx, w, err)
		return
	}

	renderer.Success(ctx, w)
}

// NewSampleHandler ... ハンドラを作成する
func NewSampleHandler(svc service.Sample) *SampleHandler {
	return &SampleHandler{
		Svc: svc,
	}
}

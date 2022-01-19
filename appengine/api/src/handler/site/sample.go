package site

import (
	"net/http"

	"github.com/rabee-inc/go-pkg/renderer"

	"github.com/aikizoku/rabbitgo/appengine/api/src/service"
)

type SampleHandler struct {
	Svc service.Sample
}

func NewSampleHandler(svc service.Sample) *SampleHandler {
	return &SampleHandler{
		Svc: svc,
	}
}

func (h *SampleHandler) Sample(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := h.Svc.Sample(ctx)
	if err != nil {
		renderer.HandleError(ctx, w, err)
		return
	}

	renderer.Success(ctx, w)
}

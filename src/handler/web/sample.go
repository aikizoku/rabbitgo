package web

import (
	"net/http"

	"github.com/aikizoku/go-gae-template/src/service"
)

type Sample struct {
	svc service.Sample
}

func (h *Sample) Hoge(w http.ResponseWriter, r *http.Request) {

}

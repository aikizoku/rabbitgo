package api

import (
	"net/http"

	"github.com/aikizoku/go-web-template/src/middleware/jsonrpc2"
	"github.com/aikizoku/go-web-template/src/service"
)

type Sample struct {
	sample service.Sample
}

func (h *Sample) Jsonrpc2(w http.ResponseWriter, r *http.Request) {
	rpc := jsonrpc2.Jsonrpc2{}
}

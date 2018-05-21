package api

import (
	"net/http"

	"github.com/aikizoku/go-gae-template/src/middleware"
	"google.golang.org/appengine"
)

type Jsonrpc2 struct {
	Rpc middleware.Jsonrpc2
}

func (h *Jsonrpc2) Handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	h.Rpc.Handle(ctx, w, r)
}

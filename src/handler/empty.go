package handler

import (
	"net/http"

	"github.com/aikizoku/merlin/src/lib/log"
)

// Empty ... 空のハンドラ
func Empty(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Debugf(ctx, "debug %d", 123)
	log.Infof(ctx, "info %d", 123)
	log.Warningf(ctx, "warning %d", 123)
	log.Errorf(ctx, "error %d", 123)
	log.Criticalf(ctx, "critical %d", 123)

	RenderSuccess(w)
}

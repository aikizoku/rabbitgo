package handler

import (
	"net/http"

	"github.com/rabee-inc/go-pkg/log"
)

// Ping ... 生存確認
func Ping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
	log.SetResponseStatus(ctx, http.StatusOK)
}

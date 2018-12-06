package log

import (
	"net/http"
	"os"
)

// Handle ... ロガーを初期化する
func Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		minLogLevel := NewLevel(os.Getenv("MIN_LOG_LEVEL"))
		logger := NewLogger(minLogLevel)
		ctx := r.Context()
		ctx = SetLogger(ctx, logger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

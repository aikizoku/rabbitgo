package log

import (
	"net/http"

	"github.com/aikizoku/rabbitgo/src/lib/util"
)

// Middleware ... ロガー
type Middleware struct {
	Writer         Writer
	MinOutSeverity Severity
}

// Handle ... ロガーを初期化する
func (m *Middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startAt := util.TimeNow()

		// ロガーをContextに設定
		traceID := util.StrUniqueID()
		logger := NewLogger(m.Writer, m.MinOutSeverity, traceID)
		ctx := r.Context()
		ctx = SetLogger(ctx, logger)

		// 実行
		next.ServeHTTP(w, r.WithContext(ctx))

		// 実行時間を計算
		endAt := util.TimeNow()
		dr := endAt.Sub(startAt)

		// リクエストログを出力
		logger.WriteRequest(r, endAt, dr)
	})
}

// NewMiddleware ... ミドルウェアを作成する
func NewMiddleware(writer Writer, minOutSeverity string) *Middleware {
	mos := NewSeverity(minOutSeverity)
	return &Middleware{
		Writer:         writer,
		MinOutSeverity: mos,
	}
}

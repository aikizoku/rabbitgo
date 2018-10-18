package log

import (
	"context"
)

type contextKey string

const loggerContextKey contextKey = "log:logger"

// GetLogger ... HTTPHeaderの値を取得
func GetLogger(ctx context.Context) Logger {
	return ctx.Value(loggerContextKey).(Logger)
}

// SetLogger ... HTTPHeaderの値を設定
func SetLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

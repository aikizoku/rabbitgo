package mysql

import (
	"context"

	"github.com/aikizoku/rabbitgo/src/lib/log"
)

// Logger ... Gorm用のカスタムロガー
type Logger struct {
	ctx context.Context
}

// Println ... StackDriverLoggingに出力する
func (l *Logger) Println(values ...interface{}) {
	texts := ""
	for _, value := range values {
		if text, ok := value.(string); ok {
			texts += text
		}
	}
	log.Infof(l.ctx, texts)
}

// NewLogger ... ロガーを作成する
func NewLogger(ctx context.Context) *Logger {
	return &Logger{
		ctx: ctx,
	}
}

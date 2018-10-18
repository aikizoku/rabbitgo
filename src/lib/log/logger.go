package log

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Logger ... ロガー
type Logger struct {
	MinLevel  Level
	Method    string
	URL       string
	UserAgent string
}

// IsLogging ... レベル毎のログ出力許可
func (l *Logger) IsLogging(level Level) bool {
	return l.MinLevel <= level
}

// NewLogger ... Loggerを作成する
func NewLogger(minLevel Level, method string, url string, ua string) Logger {
	return Logger{
		MinLevel:  minLevel,
		Method:    method,
		URL:       url,
		UserAgent: ua,
	}
}

// Debugf ... Debugログを出力する
func Debugf(ctx context.Context, format string, args ...interface{}) {
	writeLog(ctx, LevelDebug, format, args...)
}

// Infof ... Infoログを出力する
func Infof(ctx context.Context, format string, args ...interface{}) {
	writeLog(ctx, LevelInfo, format, args...)
}

// Warningf ... Warningログを出力する
func Warningf(ctx context.Context, format string, args ...interface{}) {
	writeLog(ctx, LevelWarning, format, args...)
}

// Errorf ... Errorログを出力する
func Errorf(ctx context.Context, format string, args ...interface{}) {
	writeLog(ctx, LevelError, format, args...)
}

// Criticalf ... Criticalログを出力する
func Criticalf(ctx context.Context, format string, args ...interface{}) {
	writeLog(ctx, LevelCritical, format, args...)
}

func writeLog(ctx context.Context, level Level, format string, args ...interface{}) {
	logger := GetLogger(ctx)
	if !logger.IsLogging(level) {
		return
	}

	entry := Entry{}
	entry.Method = logger.Method
	entry.URL = logger.URL
	entry.UserAgent = logger.UserAgent
	entry.Level = level
	entry.Message = fmt.Sprintf(format, args...)
	if _, file, line, ok := runtime.Caller(2); ok {
		println(line)
		parts := strings.Split(file, "/")
		length := len(parts)
		entry.File = fmt.Sprintf("%s/%s", parts[length-2], parts[length-1])
		entry.Line = line
	}

	text := fmt.Sprintf(
		"[%s] [%s %s] [%s] %s:%d %s",
		entry.Level.String(),
		entry.Method,
		entry.URL,
		entry.UserAgent,
		entry.File,
		entry.Line,
		entry.Message,
	)
	_, _ = os.Stderr.WriteString(text + "\n")
}

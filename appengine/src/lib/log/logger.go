package log

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/aikizoku/rabbitgo/appengine/src/lib/errcode"
	"github.com/aikizoku/rabbitgo/appengine/src/lib/util"
)

// Logger ... ロガー
type Logger struct {
	Writer            Writer
	MinOutSeverity    Severity
	MaxOuttedSeverity Severity
	TraceID           string
	ResponseStatus    int
}

// IsLogging ... レベル毎のログ出力許可
func (l *Logger) IsLogging(severity Severity) bool {
	return l.MinOutSeverity <= severity
}

// SetOuttedSeverity ... 出力された最大のレベルを設定
func (l *Logger) SetOuttedSeverity(severity Severity) {
	if l.MaxOuttedSeverity < severity {
		l.MaxOuttedSeverity = severity
	}
}

// WriteRequest ... リクエストログを出力する
func (l *Logger) WriteRequest(r *http.Request, at time.Time, dr time.Duration) {
	l.Writer.Request(
		l.MaxOuttedSeverity,
		l.TraceID,
		r,
		l.ResponseStatus,
		at,
		dr)
}

// NewLogger ... Loggerを作成する
func NewLogger(writer Writer, minSeverity Severity, traceID string) *Logger {
	return &Logger{
		Writer:         writer,
		MinOutSeverity: minSeverity,
		TraceID:        traceID,
	}
}

// SetResponseStatus ... レスポンスのステータスコードを設定する
func SetResponseStatus(ctx context.Context, status int) {
	logger := GetLogger(ctx)
	if logger != nil {
		logger.ResponseStatus = status
	}
}

// Debugf ... Debugログを出力する
func Debugf(ctx context.Context, format string, args ...interface{}) {
	severity := SeverityDebug
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			fmt.Sprintf(format, args...),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
}

// Debugm ... Debugログの定形を出力する
func Debugm(ctx context.Context, method string, err error) {
	severity := SeverityDebug
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			fmt.Sprintf("%s: %s", method, err.Error()),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
}

// Debuge ... Debugログを出力してエラーを生成する
func Debuge(ctx context.Context, format string, args ...interface{}) error {
	err := fmt.Errorf(format, args...)
	severity := SeverityDebug
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			err.Error(),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
	return err
}

// Debugc ... Debugログを出力してコード付きのエラーを生成する
func Debugc(ctx context.Context, code int, format string, args ...interface{}) error {
	err := fmt.Errorf(format, args...)
	severity := SeverityDebug
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			err.Error(),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
	return errcode.Set(err, code)
}

// Infof ... Infoログを出力する
func Infof(ctx context.Context, format string, args ...interface{}) {
	severity := SeverityInfo
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			fmt.Sprintf(format, args...),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
}

// Infom ... Infoログの定形を出力する
func Infom(ctx context.Context, method string, err error) {
	severity := SeverityInfo
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			fmt.Sprintf("%s: %s", method, err.Error()),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
}

// Infoe ... Infoログを出力してエラーを生成する
func Infoe(ctx context.Context, format string, args ...interface{}) error {
	err := fmt.Errorf(format, args...)
	severity := SeverityInfo
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			err.Error(),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
	return err
}

// Infoc ... Infoログを出力してコード付きのエラーを生成する
func Infoc(ctx context.Context, code int, format string, args ...interface{}) error {
	err := fmt.Errorf(format, args...)
	severity := SeverityInfo
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			err.Error(),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
	return errcode.Set(err, code)
}

// Warningf ... Warningログを出力する
func Warningf(ctx context.Context, format string, args ...interface{}) {
	severity := SeverityWarning
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			fmt.Sprintf(format, args...),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
}

// Warningm ... Warningログの定形を出力する
func Warningm(ctx context.Context, method string, err error) {
	severity := SeverityWarning
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			fmt.Sprintf("%s: %s", method, err.Error()),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
}

// Warninge ... Warningログを出力してエラーを生成する
func Warninge(ctx context.Context, format string, args ...interface{}) error {
	err := fmt.Errorf(format, args...)
	severity := SeverityWarning
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			err.Error(),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
	return err
}

// Warningc ... Warningログを出力してコード付きのエラーを生成する
func Warningc(ctx context.Context, code int, format string, args ...interface{}) error {
	err := fmt.Errorf(format, args...)
	severity := SeverityWarning
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			err.Error(),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
	return errcode.Set(err, code)
}

// Errorf ... Errorログを出力する
func Errorf(ctx context.Context, format string, args ...interface{}) {
	severity := SeverityError
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			fmt.Sprintf(format, args...),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
}

// Errorm ... Errorログの定形を出力する
func Errorm(ctx context.Context, method string, err error) {
	severity := SeverityError
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			fmt.Sprintf("%s: %s", method, err.Error()),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
}

// Errore ... Errorログを出力してエラーを生成する
func Errore(ctx context.Context, format string, args ...interface{}) error {
	err := fmt.Errorf(format, args...)
	severity := SeverityError
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			err.Error(),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
	return err
}

// Errorc ... Errorログを出力してコード付きのエラーを生成する
func Errorc(ctx context.Context, code int, format string, args ...interface{}) error {
	err := fmt.Errorf(format, args...)
	severity := SeverityError
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			err.Error(),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
	return errcode.Set(err, code)
}

// Criticalf ... Criticalログを出力する
func Criticalf(ctx context.Context, format string, args ...interface{}) {
	severity := SeverityCritical
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			fmt.Sprintf(format, args...),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
}

// Criticalm ... Criticalログの定形を出力する
func Criticalm(ctx context.Context, method string, err error) {
	severity := SeverityCritical
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			fmt.Sprintf("%s: %s", method, err.Error()),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
}

// Criticale ... Criticalログを出力してエラーを生成する
func Criticale(ctx context.Context, format string, args ...interface{}) error {
	err := fmt.Errorf(format, args...)
	severity := SeverityCritical
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			err.Error(),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
	return err
}

// Criticalc ... Criticalログを出力してコード付きのエラーを生成する
func Criticalc(ctx context.Context, code int, format string, args ...interface{}) error {
	err := fmt.Errorf(format, args...)
	severity := SeverityCritical
	logger := GetLogger(ctx)
	if logger != nil && logger.IsLogging(severity) {
		now := util.TimeNow()
		file, line, function := getFileLine()
		logger.Writer.Application(
			severity,
			logger.TraceID,
			err.Error(),
			file,
			line,
			function,
			now)
		logger.SetOuttedSeverity(severity)
	}
	return errcode.Set(err, code)
}

func getFileLine() (string, int64, string) {
	if pt, file, line, ok := runtime.Caller(2); ok {
		parts := strings.Split(file, "/")
		length := len(parts)
		file := fmt.Sprintf("%s/%s", parts[length-2], parts[length-1])

		fParts := strings.Split(runtime.FuncForPC(pt).Name(), ".")
		fLength := len(fParts)
		return file, int64(line), fParts[fLength-1]
	}
	return "", 0, ""
}

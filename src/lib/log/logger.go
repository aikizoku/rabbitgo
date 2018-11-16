package log

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"google.golang.org/appengine/log"
)

// Debugf ... Debugログを出力する
func Debugf(ctx context.Context, format string, args ...interface{}) {
	fl := getFileLine()
	log.Debugf(ctx, fl+format, args...)
}

// Infof ... Infoログを出力する
func Infof(ctx context.Context, format string, args ...interface{}) {
	fl := getFileLine()
	log.Infof(ctx, fl+format, args...)
}

// Warningf ... Warningログを出力する
func Warningf(ctx context.Context, format string, args ...interface{}) {
	fl := getFileLine()
	log.Warningf(ctx, fl+format, args...)
}

// Errorf ... Errorログを出力する
func Errorf(ctx context.Context, format string, args ...interface{}) {
	fl := getFileLine()
	log.Errorf(ctx, fl+format, args...)
}

// Criticalf ... Criticalログを出力する
func Criticalf(ctx context.Context, format string, args ...interface{}) {
	fl := getFileLine()
	log.Criticalf(ctx, fl+format, args...)
}

func getFileLine() string {
	var ret string
	if _, file, line, ok := runtime.Caller(2); ok {
		parts := strings.Split(file, "/")
		length := len(parts)
		ret = fmt.Sprintf("%s/%s:%d ", parts[length-2], parts[length-1], line)
	}
	return ret
}

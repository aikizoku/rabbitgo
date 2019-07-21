package log

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type writerStackdriver struct {
	ProjectID string
}

func (w *writerStackdriver) Request(
	severity Severity,
	traceID string,
	r *http.Request,
	status int,
	at time.Time,
	dr time.Duration) {
	u := *r.URL
	u.Fragment = ""

	remoteAddr := ""
	if v := r.Header.Get("X-AppEngine-User-IP"); v != "" {
		remoteAddr = v
	} else if v := r.Header.Get("X-Forwarded-For"); v != "" {
		remoteAddr = v
	} else {
		remoteAddr = strings.SplitN(r.RemoteAddr, ":", 2)[0]
	}

	falseV := false

	e := &Entry{
		Severity: severity.String(),
		Time:     Time(at),
		Trace:    fmt.Sprintf("projects/%s/traces/%s", w.ProjectID, traceID),
		Message:  "",
		HTTPRequest: &EntryHTTPRequest{
			RequestMethod:                  r.Method,
			RequestURL:                     u.RequestURI(),
			RequestSize:                    r.ContentLength,
			Status:                         status,
			UserAgent:                      r.UserAgent(),
			RemoteIP:                       remoteAddr,
			Referer:                        r.Referer(),
			Latency:                        Duration(dr),
			CacheLookup:                    &falseV,
			CacheHit:                       &falseV,
			CacheValidatedWithOriginServer: &falseV,
			CacheFillBytes:                 nil,
			Protocol:                       r.Proto,
		},
	}
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, string(b)+"\n")
}

func (w *writerStackdriver) Application(
	severity Severity,
	traceID string,
	msg string,
	file string,
	line int64,
	function string,
	at time.Time) {
	e := &Entry{
		Severity: severity.String(),
		Time:     Time(at),
		Trace:    fmt.Sprintf("projects/%s/traces/%s", w.ProjectID, traceID),
		Message:  msg,
		SourceLocation: &EntrySourceLocation{
			File:     file,
			Line:     line,
			Function: function,
		},
	}
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, string(b)+"\n")
}

// NewWriterStackdriver ... ログ出力を作成する
func NewWriterStackdriver(projectID string) Writer {
	return &writerStackdriver{
		ProjectID: projectID,
	}
}

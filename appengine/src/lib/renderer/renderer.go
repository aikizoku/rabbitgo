package renderer

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/unrolled/render"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/aikizoku/rabbitgo/appengine/src/lib/errcode"
	"github.com/aikizoku/rabbitgo/appengine/src/lib/log"
)

// HandleError ... 一番典型的なエラーハンドリング
func HandleError(ctx context.Context, w http.ResponseWriter, msg string, err error) {
	code, ok := errcode.Get(err)
	if !ok {
		Error(ctx, w, http.StatusInternalServerError, err.Error())
		return
	}
	switch code {
	case http.StatusBadRequest:
		msg := fmt.Sprintf("%d StatusBadRequest: %s, %s", code, msg, err.Error())
		log.Warningf(ctx, msg)
		Error(ctx, w, code, msg)
	case http.StatusForbidden:
		msg := fmt.Sprintf("%d Forbidden: %s, %s", code, msg, err.Error())
		log.Warningf(ctx, msg)
		Error(ctx, w, code, msg)
	case http.StatusNotFound:
		msg := fmt.Sprintf("%d NotFound: %s, %s", code, msg, err.Error())
		log.Warningf(ctx, msg)
		Error(ctx, w, code, msg)
	default:
		msg := fmt.Sprintf("%d: %s, %s", code, msg, err.Error())
		log.Errorf(ctx, msg)
		Error(ctx, w, code, msg)
	}
}

// Success ... 成功レスポンスをレンダリングする
func Success(ctx context.Context, w http.ResponseWriter) {
	status := http.StatusOK
	r := render.New()
	r.JSON(w, http.StatusOK, NewResponseOK(http.StatusOK))
	log.SetResponseStatus(ctx, status)
}

// Error ... エラーレスポンスをレンダリングする
func Error(ctx context.Context, w http.ResponseWriter, status int, msg string) {
	r := render.New()
	r.JSON(w, status, NewResponseError(status, msg))
	log.SetResponseStatus(ctx, status)
}

// JSON ... JSONをレンダリングする
func JSON(ctx context.Context, w http.ResponseWriter, status int, v interface{}) {
	r := render.New()
	r.JSON(w, status, v)
	log.SetResponseStatus(ctx, status)
}

// HTML ... HTMLをレンダリングする
func HTML(ctx context.Context, w http.ResponseWriter, status int, name string, values interface{}) {
	r := render.New()
	r.HTML(w, status, name, values)
	log.SetResponseStatus(ctx, status)
}

// CSV ... CSVをレンダリングする
func CSV(ctx context.Context, w http.ResponseWriter, name string, data [][]string) {
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s.csv", name))

	writer := csv.NewWriter(transform.NewWriter(w, japanese.ShiftJIS.NewEncoder()))
	for _, datum := range data {
		writer.Write(datum)
	}
	writer.Flush()
	log.SetResponseStatus(ctx, http.StatusOK)
}

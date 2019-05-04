package renderer

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/aikizoku/merlin/src/lib/errcode"
	"github.com/aikizoku/merlin/src/lib/log"
	"github.com/unrolled/render"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// HandleError ... 一番典型的なエラーハンドリング
func HandleError(ctx context.Context, w http.ResponseWriter, msg string, err error) {
	code, ok := errcode.Get(err)
	if !ok {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	switch code {
	case http.StatusBadRequest:
		msg := fmt.Sprintf("%d StatusBadRequest: %s, %s", code, msg, err.Error())
		log.Warningf(ctx, msg)
		Error(w, code, msg)
	case http.StatusForbidden:
		msg := fmt.Sprintf("%d Forbidden: %s, %s", code, msg, err.Error())
		log.Warningf(ctx, msg)
		Error(w, code, msg)
	case http.StatusNotFound:
		msg := fmt.Sprintf("%d NotFound: %s, %s", code, msg, err.Error())
		log.Warningf(ctx, msg)
		Error(w, code, msg)
	default:
		msg := fmt.Sprintf("%d: %s, %s", code, msg, err.Error())
		log.Errorf(ctx, msg)
		Error(w, code, msg)
	}
}

// Success ... 成功レスポンスをレンダリングする
func Success(w http.ResponseWriter) {
	r := render.New()
	r.JSON(w, http.StatusOK, NewResponseOK(http.StatusOK))
}

// Error ... エラーレスポンスをレンダリングする
func Error(w http.ResponseWriter, status int, msg string) {
	r := render.New()
	r.JSON(w, status, NewResponseError(status, msg))
}

// JSON ... JSONをレンダリングする
func JSON(w http.ResponseWriter, status int, v interface{}) {
	r := render.New()
	r.JSON(w, status, v)
}

// HTML ... HTMLをレンダリングする
func HTML(w http.ResponseWriter, status int, name string, values interface{}) {
	r := render.New()
	r.HTML(w, status, name, values)
}

// CSV ... CSVをレンダリングする
func CSV(w http.ResponseWriter, name string, data [][]string) {
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s.csv", name))

	writer := csv.NewWriter(transform.NewWriter(w, japanese.ShiftJIS.NewEncoder()))
	for _, datum := range data {
		writer.Write(datum)
	}
	writer.Flush()
}

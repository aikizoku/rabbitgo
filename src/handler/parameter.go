package handler

import (
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/aikizoku/merlin/src/lib/log"
	"github.com/go-chi/chi"
)

// GetURLParam ... リクエストからURLParamを取得する
func GetURLParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

// GetURLParamByInt64 ... リクエストからURLParamをint64で取得する
func GetURLParamByInt64(ctx context.Context, r *http.Request, key string) (int64, error) {
	str := chi.URLParam(r, key)
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Warningm(ctx, "strconv.ParseInt", err)
		return num, err
	}
	return num, nil
}

// GetFormValue ... リクエストからFormValueを取得する
func GetFormValue(r *http.Request, key string) string {
	return r.FormValue(key)
}

// GetFormValueByInt64 ... リクエストからFormValueをint64で取得する
func GetFormValueByInt64(ctx context.Context, r *http.Request, key string) (int64, error) {
	str := chi.URLParam(r, key)
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Warningm(ctx, "strconv.ParseInt", err)
		return num, err
	}
	return num, nil
}

// GetJSON ... リクエストからJSONを取得する
func GetJSON(r *http.Request, dst interface{}) error {
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(dst)
	if err != nil {
		ctx := r.Context()
		log.Warningm(ctx, "dec.Decode", err)
		return err
	}
	return nil
}

// GetFormFile ... リクエストからファイルを取得する
func GetFormFile(r *http.Request, key string) (multipart.File, *multipart.FileHeader, error) {
	return r.FormFile(key)
}

package parameter

import (
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/aikizoku/merlin/src/lib/log"
	"github.com/go-chi/chi"
)

// GetURL ... リクエストからURLパラメータを取得する
func GetURL(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

// GetURLByInt64 ... リクエストからURLパラメータをint64で取得する
func GetURLByInt64(ctx context.Context, r *http.Request, key string) (int64, error) {
	str := chi.URLParam(r, key)
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Warningm(ctx, "strconv.ParseInt", err)
		return num, err
	}
	return num, nil
}

// GetURLByFloat64 ... リクエストからURLパラメータをfloat64で取得する
func GetURLByFloat64(ctx context.Context, r *http.Request, key string) (float64, error) {
	str := chi.URLParam(r, key)
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Warningm(ctx, "strconv.ParseFloat", err)
		return num, err
	}
	return num, nil
}

// GetForm ... リクエストからFormパラメータを取得する
func GetForm(r *http.Request, key string) string {
	return r.FormValue(key)
}

// GetFormByInt64 ... リクエストからFormパラメータをint64で取得する
func GetFormByInt64(ctx context.Context, r *http.Request, key string) (int64, error) {
	str := r.FormValue(key)
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Warningm(ctx, "strconv.ParseInt", err)
		return num, err
	}
	return num, nil
}

// GetFormByFloat64 ... リクエストからFormパラメータをfloat64で取得する
func GetFormByFloat64(ctx context.Context, r *http.Request, key string) (float64, error) {
	str := r.FormValue(key)
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Warningm(ctx, "strconv.ParseFloat", err)
		return num, err
	}
	return num, nil
}

// GetJSON ... リクエストからJSONパラメータを取得する
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

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aikizoku/beego/src/lib/log"
	"github.com/go-chi/chi"
)

// GetURLParam ... リクエストからURLParamを取得する
func GetURLParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

// GetFormValue ... リクエストからFormValueを取得する
func GetFormValue(r *http.Request, key string) string {
	return r.FormValue(key)
}

// GetJSON ... リクエストからJSONを取得する
func GetJSON(r *http.Request, dst interface{}) error {
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(dst)
	if err != nil {
		ctx := r.Context()
		log.Errorf(ctx, "dec.Decode error: %s", err.Error())
		return err
	}
	return nil
}

package parameter

import (
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/aikizoku/rabbitgo/src/lib/log"
	"github.com/aikizoku/rabbitgo/src/lib/util"
)

// GetURL ... リクエストからURLパラメータを取得する
func GetURL(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

// GetURLByInt ... リクエストからURLパラメータをintで取得する
func GetURLByInt(ctx context.Context, r *http.Request, key string) (int, error) {
	str := chi.URLParam(r, key)
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Warningm(ctx, "strconv.Atoi", err)
		return num, err
	}
	return num, nil
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

// GetForm ... リクエストからFormパラメータをstringで取得する
func GetForm(r *http.Request, key string) string {
	return r.FormValue(key)
}

// GetFormByInt ... リクエストからFormパラメータをintで取得する
func GetFormByInt(ctx context.Context, r *http.Request, key string) (int, error) {
	str := r.FormValue(key)
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Warningm(ctx, "strconv.Atoi", err)
		return num, err
	}
	return num, nil
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

// GetFormByBool ... リクエストからFormパラメータをboolで取得する
func GetFormByBool(ctx context.Context, r *http.Request, key string) (bool, error) {
	str := r.FormValue(key)
	val, err := strconv.ParseBool(str)
	if err != nil {
		log.Warningm(ctx, "strconv.ParseInt", err)
		return val, err
	}
	return val, nil
}

// GetForms ... リクエストからFormパラメータを取得する
func GetForms(ctx context.Context, r *http.Request, dst interface{}) error {
	if reflect.TypeOf(dst).Kind() != reflect.Ptr {
		err := log.Errore(ctx, "dst isn't a pointer")
		return err
	}

	paramType := reflect.TypeOf(dst).Elem()
	paramValue := reflect.ValueOf(dst).Elem()

	fieldCount := paramType.NumField()
	for i := 0; i < fieldCount; i++ {
		field := paramType.Field(i)

		formTag := paramType.Field(i).Tag.Get("form")
		if util.IsZero(formTag) {
			continue
		}

		fieldValue := paramValue.FieldByName(field.Name)
		if !fieldValue.CanSet() {
			err := log.Warningc(ctx, http.StatusBadRequest, "fieldValue.CanSet")
			return err
		}
		switch field.Type.Kind() {
		case reflect.Int64:
			val, err := GetFormByInt64(ctx, r, formTag)
			if err != nil {
				log.Debugm(ctx, "GetFormByInt64", err)
			}
			fieldValue.SetInt(val)
		case reflect.Int:
			val, err := GetFormByInt64(ctx, r, formTag)
			if err != nil {
				log.Debugm(ctx, "GetFormByInt64", err)
			}
			fieldValue.SetInt(val)
		case reflect.Float64:
			val, err := GetFormByFloat64(ctx, r, formTag)
			if err != nil {
				log.Debugm(ctx, "GetFormByFloat64", err)
			}
			fieldValue.SetFloat(val)
		case reflect.String:
			val := GetForm(r, formTag)
			fieldValue.SetString(val)
		case reflect.Bool:
			val, err := GetFormByBool(ctx, r, formTag)
			if err != nil {
				log.Debugm(ctx, "GetFormByBool", err)
			}
			fieldValue.SetBool(val)
		}
	}
	return nil
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

package cloudfirestore

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"cloud.google.com/go/firestore"

	"github.com/aikizoku/rabbitgo/appengine/src/lib/log"
)

func setDocByDst(dst interface{}, ref *firestore.DocumentRef) {
	rv := reflect.Indirect(reflect.ValueOf(dst))
	rt := rv.Type()
	if rt.Kind() == reflect.Struct {
		for i := 0; i < rt.NumField(); i++ {
			f := rt.Field(i)
			tag := f.Tag.Get("cloudfirestore")
			if tag == "id" && f.Type.Kind() == reflect.String {
				rv.Field(i).SetString(ref.ID)
				continue
			}
			if tag == "ref" && f.Type.Kind() == reflect.Ptr {
				rv.Field(i).Set(reflect.ValueOf(ref))
				continue
			}
		}
	}
}

func setDocByDsts(rv reflect.Value, rt reflect.Type, ref *firestore.DocumentRef) {
	if rt.Kind() == reflect.Struct {
		for i := 0; i < rt.NumField(); i++ {
			f := rt.Field(i)
			tag := f.Tag.Get("cloudfirestore")
			if tag == "id" && f.Type.Kind() == reflect.String {
				rv.Elem().Field(i).SetString(ref.ID)
				continue
			}
			if tag == "ref" && f.Type.Kind() == reflect.Ptr {
				rv.Elem().Field(i).Set(reflect.ValueOf(ref))
				continue
			}
		}
	}
}

func dump(ctx context.Context, method string, path string) {
	msg := fmt.Sprintf("[firestore] [%s] %s", method, path)
	log.Debugf(ctx, msg)
}

func dumps(ctx context.Context, method string, paths []string) {
	msg := fmt.Sprintf("[firestore] [%s]\n%s", method, strings.Join(paths, ", "))
	log.Debugf(ctx, msg)
}

package headerparams

import (
	"context"
	"net/http"
)

// Service ... Headerに関する機能を提供する
type Service interface {
	Get(ctx context.Context, r *http.Request) (HeaderParams, error)
}

package jsonrpc2

import (
	"context"
	"encoding/json"
)

// Handler ... JSORPC2ハンドラの定義
type Handler interface {
	DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error)
	Exec(ctx context.Context, method string, params interface{}) (interface{}, error)
}

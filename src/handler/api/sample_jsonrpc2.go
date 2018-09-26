package api

import (
	"context"
	"encoding/json"

	"github.com/aikizoku/beego/src/model"
	"github.com/aikizoku/beego/src/service"
)

// BeegoJSONRPC2Handler ... JSONRPC2のハンドラ
type BeegoJSONRPC2Handler struct {
	Svc service.SampleService
}

type beegoJSONRPC2Params struct {
	Sample string `json:"sample"`
}

// DecodeParams ... 受け取ったJSONパラメータをデコードする
func (h *BeegoJSONRPC2Handler) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	var params beegoJSONRPC2Params
	err := json.Unmarshal(*msg, &params)
	return params, err
}

// Exec ... 処理をする
func (h *BeegoJSONRPC2Handler) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	// パラメータを取得
	sample := params.(model.Sample)

	// Serviceを実行する
	sample, err := h.Svc.Sample(ctx)
	if err != nil {
		return nil, err
	}

	return struct {
		Sample model.Sample `json:"sample"`
		Hoge   string       `json:"hoge,omitempty"`
	}{
		Sample: sample,
		Hoge:   "",
	}, nil
}

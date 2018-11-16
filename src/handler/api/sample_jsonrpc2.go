package api

import (
	"context"
	"encoding/json"

	"github.com/aikizoku/beego/src/lib/log"
	"github.com/aikizoku/beego/src/model"
	"github.com/aikizoku/beego/src/service"
)

// SampleJSONRPC2Handler ... JSONRPC2のハンドラ
type SampleJSONRPC2Handler struct {
	Svc service.Sample
}

type sampleJSONRPC2Params struct {
	Hoge string `json:"hoge"`
}

// DecodeParams ... 受け取ったJSONパラメータをデコードする
func (h *SampleJSONRPC2Handler) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	var params sampleJSONRPC2Params
	err := json.Unmarshal(*msg, &params)
	return params, err
}

// Exec ... 処理をする
func (h *SampleJSONRPC2Handler) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	// パラメータを取得
	hoge := params.(sampleJSONRPC2Params).Hoge
	log.Debugf(ctx, hoge)

	// Serviceを実行する
	sample, err := h.Svc.Sample(ctx)
	if err != nil {
		return nil, err
	}

	return struct {
		Sample model.Sample `json:"sample"`
		Foobar string       `json:"foobar,omitempty"`
	}{
		Sample: sample,
		Foobar: "",
	}, nil
}

// NewSampleJSONRPC2Handler ... SampleJSONRPC2Handlerを作成する
func NewSampleJSONRPC2Handler(svc service.Sample) *SampleJSONRPC2Handler {
	return &SampleJSONRPC2Handler{
		Svc: svc,
	}
}

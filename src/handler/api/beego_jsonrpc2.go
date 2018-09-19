package api

import (
	"context"
	"encoding/json"

	"github.com/aikizoku/beego/src/model"
	"github.com/aikizoku/beego/src/service"
)

// BeegoJSONRPC2Handler ... JSONRPC2のハンドラ
type BeegoJSONRPC2Handler struct {
	Beego service.Beego
}

type beegoJSONRPC2Params struct {
	Beego string `json:"beego"`
}

// DecodeParams ... 受け取ったJSONパラメータをデコードする
func (s *BeegoJSONRPC2Handler) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	var params beegoJSONRPC2Params
	err := json.Unmarshal(*msg, &params)
	return params, err
}

// Exec ... 処理をする
func (s *BeegoJSONRPC2Handler) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	// パラメータを取得
	beego := params.(model.Beego)

	// Serviceを実行する
	beego, err := h.Service.Beego(ctx)
	if err != nil {
		return nil, err
	}

	return struct {
		Beego model.Beego `json:"beego"`
		Hoge  string      `json:"hoge,omitempty"`
	}{
		Beego: beego,
		Hoge:  "",
	}, nil
}

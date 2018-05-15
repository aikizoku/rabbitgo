package jsonrpc2

import "encoding/json"

// Request ... JSONRPC2のリクエスト定義
type Request struct {
	Version string           `json:"jsonrpc"`
	ID      string           `json:"id"`
	Method  string           `json:"method"`
	Params  *json.RawMessage `json:"params"`
}

// NewRequest ... JSONRPC2のリクエストを取得
func NewRequest(id string, method string, params *json.RawMessage) Request {
	return Request{
		Version: "2.0",
		ID:      id,
		Method:  method,
		Params:  params,
	}
}

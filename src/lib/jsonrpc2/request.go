package jsonrpc2

import "encoding/json"

// ClientRequest ... JSONRPC2実行のリクエスト
type ClientRequest struct {
	Version string           `json:"jsonrpc"`
	ID      string           `json:"id"`
	Method  string           `json:"method"`
	Params  *json.RawMessage `json:"params"`
}

type request struct {
	Version string           `json:"jsonrpc"`
	ID      string           `json:"id"`
	Method  string           `json:"method"`
	Params  *json.RawMessage `json:"params"`
}

func (r *request) isValid() bool {
	if r.Version != version {
		return false
	}
	if r.ID == "" {
		return false
	}
	if r.Method == "" {
		return false
	}
	if r.Params == nil {
		return false
	}
	return true
}

// GenerateRequestBody ... JSONRPC2のリクエストBodyを作成する
func GenerateRequestBody(id string, method string, params interface{}) (*json.RawMessage, error) {
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	rawParams := json.RawMessage(jsonParams)
	request := request{
		Version: version,
		ID:      id,
		Method:  method,
		Params:  &rawParams,
	}
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	rawRequest := json.RawMessage(jsonRequest)
	return &rawRequest, nil
}

// GenerateBatchRequestBody ... JSONRPC2のバッチリクエストBodyを作成する
func GenerateBatchRequestBody(requests []*json.RawMessage) (*json.RawMessage, error) {
	jsonRequests, err := json.Marshal(requests)
	if err != nil {
		return nil, err
	}
	rawRequests := json.RawMessage(jsonRequests)
	return &rawRequests, nil
}

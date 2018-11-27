package jsonrpc2

import "encoding/json"

// ClientResponse ... JSONRPC2実行のレスポンス
type ClientResponse struct {
	Version string           `json:"jsonrpc"`
	ID      string           `json:"id"`
	Result  *json.RawMessage `json:"result,omitempty"`
	Error   *ErrorResponse   `json:"error,omitempty"`
}

type response struct {
	Version string         `json:"jsonrpc"`
	ID      string         `json:"id"`
	Result  interface{}    `json:"result,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
}

// ErrorResponse ... JSONRPC2のエラーレスポンス
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func newResponse(id string, result interface{}) response {
	return response{
		Version: version,
		ID:      id,
		Result:  result,
	}
}

func newErrorResponse(id string, code int, message string) response {
	return response{
		Version: version,
		ID:      id,
		Error: &ErrorResponse{
			Code:    code,
			Message: message,
		},
	}
}

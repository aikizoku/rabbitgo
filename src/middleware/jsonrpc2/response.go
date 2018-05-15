package jsonrpc2

// Response ... JSONRPC2のレスポンス定義
type Response struct {
	Version string        `json:"jsonrpc"`
	ID      string        `json:"id"`
	Result  interface{}   `json:"result,omitempty"`
	Error   ResponseError `json:"error,omitempty"`
}

// ResponseError ... JSONRPC2のエラーレスポンス定義
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewResponse ... JSONRPC2のレスポンスを取得
func NewResponse(id string, result interface{}) Response {
	return Response{
		Version: "2.0",
		ID:      id,
		Result:  result,
	}
}

// NewErrorResponse ... JSONRPC2のエラーレスポンスを取得
func NewErrorResponse(id string, code int, message string) Response {
	return Response{
		Version: "2.0",
		ID:      id,
		Error: ResponseError{
			Code:    code,
			Message: message,
		},
	}
}

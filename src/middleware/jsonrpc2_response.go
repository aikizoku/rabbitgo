package middleware

type jsonrpc2Response struct {
	Version string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type jsonrpc2ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func newJsonrpc2Response(id string, result interface{}) jsonrpc2Response {
	return jsonrpc2Response{
		Version: jsonrpc2Version,
		ID:      id,
		Result:  result,
	}
}

func newJsonrpc2ErrorResponse(id string, code int, message string) jsonrpc2Response {
	return jsonrpc2Response{
		Version: jsonrpc2Version,
		ID:      id,
		Error: jsonrpc2ErrorResponse{
			Code:    code,
			Message: message,
		},
	}
}

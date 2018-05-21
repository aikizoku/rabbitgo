package jsonrpc2

type response struct {
	Version string        `json:"jsonrpc"`
	ID      string        `json:"id"`
	Result  interface{}   `json:"result,omitempty"`
	Error   responseError `json:"error,omitempty"`
}

type responseError struct {
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
		Error: responseError{
			Code:    code,
			Message: message,
		},
	}
}

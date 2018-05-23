package middleware

import "encoding/json"

type jsonrpc2Request struct {
	Version string           `json:"jsonrpc"`
	ID      string           `json:"id"`
	Method  string           `json:"method"`
	Params  *json.RawMessage `json:"params"`
}

func (r *jsonrpc2Request) isValid() bool {
	if r.Version != jsonrpc2Version {
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

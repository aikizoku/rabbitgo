package jsonrpc2

import "encoding/json"

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

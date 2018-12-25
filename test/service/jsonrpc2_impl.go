package service

type jsonrpc2 struct {
}

// NewJSONRPC2 ... JSONRPC2を作成する
func NewJSONRPC2() JSONRPC2 {
	return &jsonrpc2{}
}

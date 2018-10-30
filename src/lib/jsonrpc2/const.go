package jsonrpc2

const (
	// ErrInvalidRequest ... 不正なリクエスト
	ErrInvalidRequest = 40001
	// ErrInvalidJSON ... 不正なJSON形式
	ErrInvalidJSON = 40002
	// ErrInvalidJsonrpc2 ... JSONがJSONRPC2の形式ではない
	ErrInvalidJsonrpc2 = 40003
	// ErrInvalidParams ... 不正なパラメータ
	ErrInvalidParams = 40004
	// ErrMehodNotFaund ... 存在しないMethod
	ErrMehodNotFaund = 40401
	// ErrInternal ... 内部エラー
	ErrInternal = 50001

	contentType = "application/json"
	version     = "2.0"
)

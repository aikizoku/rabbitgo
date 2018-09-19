package repository

import (
	"context"

	"github.com/aikizoku/beego/src/model"
)

// Beego ... リポジトリのインターフェースを定義
type Beego interface {
	// DataStore
	DataStoreGet(ctx context.Context, id int64) (model.Beego, error)
	DataStoreGetMulti(ctx context.Context, ids []int64) ([]model.Beego, error)
	DataStorePut(ctx context.Context, obj model.Beego) error
	DataStoreDelete(ctx context.Context, id int64) error

	// CloudSQL
	CloudSQLGet(ctx context.Context, id int64) (model.Beego, error)
	CloudSQLGetMulti(ctx context.Context, ids []int64) ([]model.Beego, error)
	CloudSQLPut(ctx context.Context, obj model.Beego) error
	CloudSQLDelete(ctx context.Context, id int64) error

	// HTTP
	HTTPGet(ctx context.Context) error
	HTTPPost(ctx context.Context) error
}

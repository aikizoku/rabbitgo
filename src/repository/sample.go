package repository

import (
	"context"

	"github.com/aikizoku/beego/src/model"
)

// SampleRepository ... リポジトリのインターフェース
type SampleRepository interface {
	// DataStore
	DataStoreGet(ctx context.Context, id int64) (model.Sample, error)
	DataStoreGetMulti(ctx context.Context, ids []int64) ([]model.Sample, error)
	DataStorePut(ctx context.Context, obj model.Sample) error
	DataStoreDelete(ctx context.Context, id int64) error

	// CloudSQL
	CloudSQLGet(ctx context.Context, id int64) (model.Sample, error)
	CloudSQLGetMulti(ctx context.Context, ids []int64) ([]model.Sample, error)
	CloudSQLInsert(ctx context.Context, obj model.Sample) error
	CloudSQLUpdate(ctx context.Context, obj model.Sample) error
	CloudSQLUpsert(ctx context.Context, obj model.Sample) error
	CloudSQLDelete(ctx context.Context, id int64) error

	// HTTP
	HTTPGet(ctx context.Context) error
	HTTPPost(ctx context.Context) error
}

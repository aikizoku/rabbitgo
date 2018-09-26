package repository

import (
	"context"
	"database/sql"

	"github.com/aikizoku/beego/src/infrastructure"
	"github.com/aikizoku/beego/src/model"
)

type sampleRepository struct {
	http *infrastructure.HTTP
	csql *sql.DB
}

// DataStore
func (r *sampleRepository) DataStoreGet(ctx context.Context, id int64) (model.Sample, error) {
	return model.Sample{}, nil
}

func (r *sampleRepository) DataStoreGetMulti(ctx context.Context, ids []int64) ([]model.Sample, error) {
	return []model.Sample{}, nil
}

func (r *sampleRepository) DataStorePut(ctx context.Context, obj model.Sample) error {
	return nil
}

func (r *sampleRepository) DataStoreDelete(ctx context.Context, id int64) error {
	return nil
}

// CloudSQL
func (r *sampleRepository) CloudSQLGet(ctx context.Context, id int64) (model.Sample, error) {
	return model.Sample{}, nil
}

func (r *sampleRepository) CloudSQLGetMulti(ctx context.Context, ids []int64) ([]model.Sample, error) {
	return []model.Sample{}, nil
}

func (r *sampleRepository) CloudSQLPut(ctx context.Context, obj model.Sample) error {
	return nil
}

func (r *sampleRepository) CloudSQLDelete(ctx context.Context, id int64) error {
	return nil
}

// HTTP
func (r *sampleRepository) HTTPGet(ctx context.Context) error {
	return nil
}

func (r *sampleRepository) HTTPPost(ctx context.Context) error {
	return nil
}

// NewSampleRepository ... サンプルリポジトリを取得する
func NewSampleRepository(http *infrastructure.HTTP, csql *sql.DB) SampleRepository {
	return &sampleRepository{
		http: http,
		csql: csql,
	}
}

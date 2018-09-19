package repository

import (
	"context"
	"database/sql"

	"github.com/aikizoku/beego/src/infrastructure"
	"github.com/aikizoku/beego/src/model"
)

type beego struct {
	http *infrastructure.HTTP
	csql *sql.DB
}

// DataStore
func (r *beego) DataStoreGet(ctx context.Context, id int64) (model.Beego, error) {
	return model.Beego{}, nil
}

func (r *beego) DataStoreGetMulti(ctx context.Context, ids []int64) ([]model.Beego, error) {
	return []model.Beego{}, nil
}

func (r *beego) DataStorePut(ctx context.Context, obj model.Beego) error {
	return nil
}

func (r *beego) DataStoreDelete(ctx context.Context, id int64) error {
	return nil
}

// CloudSQL
func (r *beego) CloudSQLGet(ctx context.Context, id int64) (model.Beego, error) {
	return model.Beego{}, nil
}

func (r *beego) CloudSQLGetMulti(ctx context.Context, ids []int64) ([]model.Beego, error) {
	return []model.Beego{}, nil
}

func (r *beego) CloudSQLPut(ctx context.Context, obj model.Beego) error {
	return nil
}

func (r *beego) CloudSQLDelete(ctx context.Context, id int64) error {
	return nil
}

// HTTP
func (r *beego) HTTPGet(ctx context.Context) error {
	return nil
}

func (r *beego) HTTPPost(ctx context.Context) error {
	return nil
}

// NewBeego ... リポジトリを取得する
func NewBeego(http *infrastructure.HTTP, csql *sql.DB) Beego {
	return &beego{
		http: http,
		csql: csql,
	}
}

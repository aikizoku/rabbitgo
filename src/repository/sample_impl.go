package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/aikizoku/beego/src/lib/cloudsql"
	"github.com/aikizoku/beego/src/lib/httpclient"
	"github.com/aikizoku/beego/src/lib/util"
	"github.com/aikizoku/beego/src/model"
	"github.com/aikizoku/beego/src/lib/log"
)

type sample struct {
	csql *sql.DB
}

// DataStore
func (r *sample) DataStoreGet(ctx context.Context, id int64) (model.Sample, error) {
	return model.Sample{}, nil
}

func (r *sample) DataStoreGetMulti(ctx context.Context, ids []int64) ([]model.Sample, error) {
	return []model.Sample{}, nil
}

func (r *sample) DataStorePut(ctx context.Context, obj model.Sample) error {
	return nil
}

func (r *sample) DataStoreDelete(ctx context.Context, id int64) error {
	return nil
}

// CloudSQL
func (r *sample) CloudSQLGet(ctx context.Context, id int64) (model.Sample, error) {
	var ret model.Sample

	q := sq.Select(
		"id",
		"name",
		"enabled",
		"created_at",
		"updated_at").
		From("sample").
		Where(sq.Eq{
			"id":      id,
			"enabled": 1,
		})

	cloudsql.DumpSelectQuery(ctx, q)

	row := q.RunWith(r.csql).QueryRowContext(ctx)
	err := row.Scan(
		&ret.ID,
		&ret.Name,
		&ret.Enabled,
		&ret.CreatedAt,
		&ret.UpdatedAt)
	if err != nil {
		log.Errorf(ctx, "CloudSQLGet: "+err.Error())
		return ret, err
	}

	return ret, nil
}

func (r *sample) CloudSQLGetMulti(ctx context.Context, ids []int64) ([]model.Sample, error) {
	var rets []model.Sample

	q := sq.Select(
		"id",
		"name",
		"enabled",
		"created_at",
		"updated_at").
		From("sample").
		Where(sq.Eq{
			"id":      ids,
			"enabled": 1,
		})

	cloudsql.DumpSelectQuery(ctx, q)

	rows, err := q.RunWith(r.csql).QueryContext(ctx)
	if err != nil {
		log.Errorf(ctx, "CloudSQLGetMulti: "+err.Error())
		return rets, err
	}

	for rows.Next() {
		var ret model.Sample
		err := rows.Scan(
			&ret.ID,
			&ret.Name,
			&ret.Enabled,
			&ret.CreatedAt,
			&ret.UpdatedAt)
		if err != nil {
			log.Errorf(ctx, "CloudSQLGet: "+err.Error())
			rows.Close()
			return rets, err
		}
		rets = append(rets, ret)
	}

	return rets, nil
}

func (r *sample) CloudSQLInsert(ctx context.Context, obj model.Sample) error {
	now := util.TimeNow()

	q := sq.Insert("sample").
		Columns("id", "name", "enabled", "created_at", "updated_at").
		Values(obj.ID, obj.Name, 1, now, now)

	cloudsql.DumpInsertQuery(ctx, q)

	_, err := q.RunWith(r.csql).ExecContext(ctx)
	if err != nil {
		log.Errorf(ctx, "CloudSQLInsert: "+err.Error())
		return err
	}

	return nil
}

func (r *sample) CloudSQLUpdate(ctx context.Context, obj model.Sample) error {
	now := util.TimeNow()

	q := sq.Update("sample").
		Set("name", obj.Name).
		Set("enabled", obj.Enabled).
		Set("updated_at", now).
		Where(sq.Eq{"id": obj.ID})

	cloudsql.DumpUpdateQuery(ctx, q)

	res, err := q.RunWith(r.csql).ExecContext(ctx)
	if err != nil {
		log.Errorf(ctx, "CloudSQLUpdate: "+err.Error())
		return err
	}

	if affected, _ := res.RowsAffected(); affected == 0 {
		err = fmt.Errorf("no affected id = %d", obj.ID)
		log.Errorf(ctx, "CloudSQLUpdate: "+err.Error())
		return err
	}

	return nil
}

func (r *sample) CloudSQLUpsert(ctx context.Context, obj model.Sample) error {
	now := util.TimeNow()

	q := sq.Insert("sample").
		Columns("id", "name", "enabled", "created_at", "updated_at").
		Values(obj.ID, obj.Name, 1, now, now).
		Suffix("ON DUPLICATE KEY UPDATE name = VALUES(name), updated_at = VALUES(updated_at)")

	cloudsql.DumpInsertQuery(ctx, q)

	_, err := q.RunWith(r.csql).ExecContext(ctx)
	if err != nil {
		log.Errorf(ctx, "CloudSQLUpsert: "+err.Error())
		return err
	}

	return nil
}

func (r *sample) CloudSQLDelete(ctx context.Context, id int64) error {
	q := sq.Delete("sample").Where(sq.Eq{"id": id})

	cloudsql.DumpDeleteQuery(ctx, q)

	res, err := q.RunWith(r.csql).ExecContext(ctx)
	if err != nil {
		log.Errorf(ctx, "CloudSQLDelete: "+err.Error())
		return err
	}

	if affected, _ := res.RowsAffected(); affected == 0 {
		err = fmt.Errorf("no affected id = %d", id)
		log.Errorf(ctx, "CloudSQLDelete: "+err.Error())
		return err
	}

	return nil
}

// HTTP
func (r *sample) HTTPGet(ctx context.Context) error {
	status, body, err := httpclient.Get(ctx, "https://www.google.co.jp/", nil)
	if err != nil {
		log.Errorf(ctx, "HTTPGet: "+err.Error())
		return err
	}
	if status != http.StatusOK {
		err := fmt.Errorf("http status: %d", status)
		return err
	}
	str := util.BytesToStr(body)
	log.Debugf(ctx, "body length: %d", len(str))
	return nil
}

func (r *sample) HTTPPost(ctx context.Context) error {
	return nil
}

// NewSample ... サンプルリポジトリを取得する
func NewSample(csql *sql.DB) Sample {
	return &sample{
		csql: csql,
	}
}

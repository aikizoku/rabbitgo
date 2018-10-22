package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/aikizoku/beego/src/lib/cloudsql"
	"github.com/aikizoku/beego/src/lib/httpclient"
	"github.com/aikizoku/beego/src/lib/log"
	"github.com/aikizoku/beego/src/lib/util"
	"github.com/aikizoku/beego/src/model"
	"github.com/mjibson/goon"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type sample struct {
	csql *sql.DB
}

// DataStore
func (r *sample) DataStoreGet(ctx context.Context, id int64) (model.Sample, error) {
	dst := model.Sample{
		ID: id,
	}
	g := goon.FromContext(ctx)
	if err := g.Get(&dst); err != nil {
		if err == datastore.ErrNoSuchEntity {
			return model.Sample{}, err
		}
		log.Errorf(ctx, "sample get error: %s", err.Error())
		return model.Sample{}, err
	}
	return dst, nil
}

func (r *sample) DataStoreGetMulti(ctx context.Context, ids []int64) ([]model.Sample, error) {
	ret := []model.Sample{}
	g := goon.FromContext(ctx)
	dsts := []model.Sample{}
	for _, id := range ids {
		dsts = append(dsts, model.Sample{
			ID: id,
		})
	}
	if err := g.GetMulti(&dsts); err != nil {
		mErr, ok := err.(appengine.MultiError)
		if !ok {
			log.Errorf(ctx, "sample get multi error: %s", err.Error())
			return ret, err
		}
		for i, err := range mErr {
			if err == nil {
				ret = append(ret, dsts[i])
				continue
			}
			if err == datastore.ErrNoSuchEntity {
				continue
			}
			log.Errorf(ctx, "sample get multi error: %s, id: %d", err.Error(), ids[i])
			return ret, err
		}
	} else {
		ret = dsts
	}
	return ret, nil
}

func (r *sample) DataStoreGetByQuery(ctx context.Context, category string) ([]model.Sample, error) {
	return []model.Sample{}, nil
}

func (r *sample) DataStoreInsert(ctx context.Context, obj model.Sample) (int64, error) {
	return 0, nil
}

func (r *sample) DataStoreInsertMulti(ctx context.Context, objs []model.Sample) ([]int64, error) {
	return []int64{}, nil
}

func (r *sample) DataStoreUpdate(ctx context.Context, obj model.Sample) (int64, error) {
	return 0, nil
}

func (r *sample) DataStoreUpdateMulti(ctx context.Context, objs []model.Sample) ([]int64, error) {
	return []int64{}, nil
}

func (r *sample) DataStoreUpsert(ctx context.Context, src model.Sample) (int64, error) {
	var id int64
	g := goon.FromContext(ctx)
	key, err := g.Put(&src)
	if err != nil {
		log.Errorf(ctx, "sample upsert datastore error: %s", err.Error())
		return id, err
	}
	id = key.IntID()
	return id, nil
}

func (r *sample) DataStoreUpsertMulti(ctx context.Context, srcs []model.Sample) ([]int64, error) {
	ids := []int64{}
	g := goon.FromContext(ctx)
	keys, err := g.PutMulti(srcs)
	if err != nil {
		mErr, ok := err.(appengine.MultiError)
		if !ok {
			log.Errorf(ctx, "sample upsert multi error: %s", err.Error())
			return ids, err
		}
		for i, err := range mErr {
			if err == nil {
				ids = append(ids, keys[i].IntID())
				continue
			}
			log.Errorf(ctx, "sample upsert multi error: %s, src: %v", err.Error(), srcs[i])
		}
	} else {
		for _, key := range keys {
			ids = append(ids, key.IntID())
		}
	}
	return ids, nil
}

func (r *sample) DataStoreDelete(ctx context.Context, id int64) (int64, error) {
	return 0, nil
}

func (r *sample) DataStoreDeleteMulti(ctx context.Context, id int64) ([]int64, error) {
	return nil, nil
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
		log.Errorf(ctx, "CloudSQLGet: %s", err.Error())
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
		log.Errorf(ctx, "CloudSQLGetMulti: %s", err.Error())
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
			log.Errorf(ctx, "CloudSQLGet: %s", err.Error())
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
		log.Errorf(ctx, "CloudSQLInsert: %s", err.Error())
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
		log.Errorf(ctx, "CloudSQLUpdate: %s", err.Error())
		return err
	}

	if affected, _ := res.RowsAffected(); affected == 0 {
		err = fmt.Errorf("no affected id = %d", obj.ID)
		log.Errorf(ctx, "CloudSQLUpdate: %s", err.Error())
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
		log.Errorf(ctx, "CloudSQLUpsert: %s", err.Error())
		return err
	}

	return nil
}

func (r *sample) CloudSQLDelete(ctx context.Context, id int64) error {
	q := sq.Delete("sample").Where(sq.Eq{"id": id})

	cloudsql.DumpDeleteQuery(ctx, q)

	res, err := q.RunWith(r.csql).ExecContext(ctx)
	if err != nil {
		log.Errorf(ctx, "CloudSQLDelete: %s", err.Error())
		return err
	}

	if affected, _ := res.RowsAffected(); affected == 0 {
		err = fmt.Errorf("no affected id = %d", id)
		log.Errorf(ctx, "CloudSQLDelete: %s", err.Error())
		return err
	}

	return nil
}

// HTTP
func (r *sample) HTTPGet(ctx context.Context) error {
	status, body, err := httpclient.Get(ctx, "https://www.google.co.jp/", nil)
	if err != nil {
		log.Errorf(ctx, "HTTPGet: %s", err.Error())
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

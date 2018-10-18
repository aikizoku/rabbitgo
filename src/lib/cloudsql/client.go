package cloudsql

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/aikizoku/beego/src/lib/log"
	_ "github.com/go-sql-driver/mysql" // Driverの読み込み
)

// NewCSQLClient ... CloudSQLのクライアントを取得する
func NewCSQLClient(cfg *CSQLConfig) *sql.DB {
	ds := fmt.Sprintf("%s:%s@cloudsql(%s)/",
		cfg.User,
		cfg.Password,
		cfg.ConnectionName)

	cli, err := sql.Open("mysql", ds)
	if err != nil {
		panic(err)
	}

	return cli
}

// DumpSelectQuery ... SELECTクエリを出力
func DumpSelectQuery(ctx context.Context, query sq.SelectBuilder) {
	qs, args, err := query.ToSql()
	if err != nil {
		log.Errorf(ctx, "DumpSelectQuery error: "+err.Error())
		return
	}
	dumpQuery(ctx, qs, args)
}

// DumpInsertQuery ... INSERTクエリを出力
func DumpInsertQuery(ctx context.Context, query sq.InsertBuilder) {
	qs, args, err := query.ToSql()
	if err != nil {
		log.Errorf(ctx, "DumpInsertQuery error: "+err.Error())
		return
	}
	dumpQuery(ctx, qs, args)
}

// DumpUpdateQuery ... UPDATEクエリを出力
func DumpUpdateQuery(ctx context.Context, query sq.UpdateBuilder) {
	qs, args, err := query.ToSql()
	if err != nil {
		log.Errorf(ctx, "DumpUpdateQuery error: "+err.Error())
		return
	}
	dumpQuery(ctx, qs, args)
}

// DumpDeleteQuery ... DELETEクエリを出力
func DumpDeleteQuery(ctx context.Context, query sq.DeleteBuilder) {
	qs, args, err := query.ToSql()
	if err != nil {
		log.Errorf(ctx, "DumpDeleteQuery error: "+err.Error())
		return
	}
	dumpQuery(ctx, qs, args)
}

func dumpQuery(ctx context.Context, queryString string, args []interface{}) {
	msg := fmt.Sprintf("[SQL Dump] %s, %s", queryString, args)
	log.Debugf(ctx, msg)
}

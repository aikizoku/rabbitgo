package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/aikizoku/beego/src/config"
	_ "github.com/go-sql-driver/mysql" // Driverの読み込み
)

// NewCSQLClient ... CloudSQLのクライアントを取得する
func NewCSQLClient(cfg *config.CSQLConfig) *sql.DB {
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

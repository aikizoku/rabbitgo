package infrastructure

import (
	"database/sql"
	"fmt"
)

type cloudSQL struct {
	ConnectionName string
	User           string
	Password       string
}

func (csc cloudSQL) Datasource() string {
	return fmt.Sprintf("%s:%s@cloudsql(%s)/",
		csc.User,
		csc.Password,
		csc.ConnectionName)
}

func NewDBClient(cName string, user string, password string) *sql.DB {
	cSQL := cloudSQL{
		ConnectionName: cName,
		User:           user,
		Password:       password,
	}
	c, err := sql.Open("mysql", cSQL.Datasource())
	if err != nil {
		panic(err)
	}
	return c
}

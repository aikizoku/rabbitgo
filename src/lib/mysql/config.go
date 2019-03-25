package mysql

import (
	"fmt"
	"os"
	"strings"
)

// Config ... 接続情報
type Config struct {
	Host     string
	User     string
	Password string
	DB       string
}

// NewConfig ... 接続情報を作成する
func NewConfig(db string) *Config {
	db = strings.ToUpper(db)

	hkey := fmt.Sprintf("MYSQL_%s_HOST", db)
	h := os.Getenv(hkey)
	if h == "" {
		panic(fmt.Errorf("no config key %s", hkey))
	}

	uKey := fmt.Sprintf("MYSQL_%s_USER", db)
	u := os.Getenv(uKey)
	if u == "" {
		panic(fmt.Errorf("no config key %s", uKey))
	}

	pKey := fmt.Sprintf("MYSQL_%s_PASSWORD", db)
	p := os.Getenv(pKey)
	if p == "" {
		panic(fmt.Errorf("no config key %s", pKey))
	}

	dKey := fmt.Sprintf("MYSQL_%s_DB", db)
	d := os.Getenv(dKey)
	if d == "" {
		panic(fmt.Errorf("no config key %s", dKey))
	}

	return &Config{
		Host:     h,
		User:     u,
		Password: p,
		DB:       d,
	}
}

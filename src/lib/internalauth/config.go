package internalauth

import (
	"fmt"
	"os"
)

const (
	authKey = "INTERNAL_AUTH"
)

// GetAuthKey ... 内部認証のKeyを取得する
func GetAuthKey() string {
	k := os.Getenv(authKey)
	if k == "" {
		panic(fmt.Errorf("no auth key internal auth: %s", authKey))
	}
	return k
}

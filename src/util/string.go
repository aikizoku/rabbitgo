package util

import (
	"crypto/sha256"
	"encoding/hex"
)

// StringToHash ... 文字列のハッシュ値を取得する
func StringToHash(str string) string {
	c := sha256.Sum256([]byte(str))
	return hex.EncodeToString(c[:])
}

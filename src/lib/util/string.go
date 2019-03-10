package util

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"unsafe"
)

const (
	letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	letterIdxMask = 0x3F
)

// StrToMD5 ... 文字列のハッシュ(MD5)を取得する
func StrToMD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// StrToSHA256 ... 文字列のハッシュ(SHA256)を取得する
func StrToSHA256(str string) string {
	c := sha256.Sum256([]byte(str))
	return hex.EncodeToString(c[:])
}

// StrToBytes ... 文字列をバイト列に変換する
func StrToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

// StrRand ... nビットのランダムな文字列を生成する。
func StrRand(n int) (string, error) {
	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	for i := 0; i < n; {
		idx := int(buf[i] & letterIdxMask)
		if idx < len(letters) {
			buf[i] = letters[idx]
			i++
		} else {
			if _, err := rand.Read(buf[i : i+1]); err != nil {
				return "", err
			}
		}
	}
	return string(buf), nil
}

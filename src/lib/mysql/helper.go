package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

const (
	gormInsertOption = "gorm:insert_option"
)

// HandleErrors ... GormのErrorsをハンドリングする
func HandleErrors(db *gorm.DB) error {
	errs := db.GetErrors()
	if len(errs) > 0 {
		msgs := []string{}
		for _, err := range errs {
			msgs = append(msgs, err.Error())
		}
		msg := strings.Join(msgs, ", ")
		return errors.New(msg)
	}
	return nil
}

// UpsertOption ... ON DUPLICATE KEY UPDATE句を作成する
func UpsertOption(columns []string) (string, string) {
	keys := []string{}
	for _, column := range columns {
		key := fmt.Sprintf("%s = VALUES(%s)", column, column)
		keys = append(keys, key)
	}
	opt := "ON DUPLICATE KEY UPDATE " + strings.Join(keys, ", ")
	return gormInsertOption, opt
}

// ToNullString ... 文字列を文字列(Nullあり)に変換する
func ToNullString(str string) sql.NullString {
	var dst sql.NullString
	if str == "" {
		dst = sql.NullString{
			Valid: false,
		}
	} else {
		dst = sql.NullString{
			String: str,
			Valid:  true,
		}
	}
	return dst
}

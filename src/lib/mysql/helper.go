package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
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

// Upsert ... アップサートを行う
func Upsert(
	db *gorm.DB,
	table string,
	src interface{},
	updateColumns []string) *gorm.DB {
	opt := generateUpsertOption(updateColumns)
	return db.Table(table).
		Set(gormInsertOption, opt).
		Create(src)
}

// BulkInsert ... バルクインサートを行う
func BulkInsert(
	db *gorm.DB,
	table string,
	columns []string,
	valuesList [][]interface{}) *gorm.DB {
	q := generateBulkInsertQuery(table, columns, valuesList)
	return db.Exec(q)
}

// BulkUpsert ... バルクアップサートを行う
func BulkUpsert(
	db *gorm.DB,
	table string,
	columns []string,
	valuesList [][]interface{},
	updateColumns []string) *gorm.DB {
	opt := generateUpsertOption(updateColumns)
	q := fmt.Sprintf("%s %s", generateBulkInsertQuery(table, columns, valuesList), opt)
	return db.Exec(q)
}

func generateUpsertOption(columns []string) string {
	keys := []string{}
	for _, column := range columns {
		key := fmt.Sprintf("%s = VALUES(%s)", column, column)
		keys = append(keys, key)
	}
	return "ON DUPLICATE KEY UPDATE " + strings.Join(keys, ", ")
}

func generateBulkInsertQuery(table string, columns []string, valuesList [][]interface{}) string {
	col := "`" + strings.Join(columns, "`, `") + "`"
	vals := []string{}
	for _, values := range valuesList {
		val := []string{}
		for _, value := range values {
			kind := reflect.Indirect(reflect.ValueOf(value)).Type().Kind()
			switch kind {
			case
				reflect.Bool,
				reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
				reflect.Float32, reflect.Float64,
				reflect.Complex64, reflect.Complex128:
				val = append(val, fmt.Sprintf("%v", value))
			case reflect.String:
				val = append(val, fmt.Sprintf("'%v'", value))
			default:
				if v, ok := value.(sql.NullString); ok {
					if v.Valid {
						val = append(val, fmt.Sprintf("'%s'", v.String))
					} else {
						val = append(val, "null")
					}
				} else if v, ok := value.(sql.NullInt64); ok {
					if v.Valid {
						val = append(val, fmt.Sprintf("%d", v.Int64))
					} else {
						val = append(val, "null")
					}
				} else if v, ok := value.(sql.NullFloat64); ok {
					if v.Valid {
						val = append(val, fmt.Sprintf("%f", v.Float64))
					} else {
						val = append(val, "null")
					}
				} else if v, ok := value.(sql.NullBool); ok {
					if v.Valid {
						if v.Bool {
							val = append(val, "true")
						} else {
							val = append(val, "false")
						}
					} else {
						val = append(val, "null")
					}
				}
			}
		}
		vals = append(vals, "("+strings.Join(val, ", ")+")")
	}
	return fmt.Sprintf("INSERT INTO `%s` (%s) VALUES %s", table, col, strings.Join(vals, ", "))
}

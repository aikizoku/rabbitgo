package validation

import (
	"strconv"
	"unicode/utf8"

	"gopkg.in/go-playground/validator.v9"
)

// MinimumString ... 必要文字数を指定(ゼロ値許容)
func MinimumString(fl validator.FieldLevel) bool {
	min, err := strconv.Atoi(fl.Param())
	if err != nil {
		return false
	}

	keyword := fl.Field().String()
	length := utf8.RuneCountInString(keyword)
	if keyword == "" || min <= length {
		return true
	}

	return false
}

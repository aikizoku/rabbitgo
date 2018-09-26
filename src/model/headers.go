package model

// HeaderParams ... リクエストヘッダーで受け取る値
type HeaderParams struct {
	Sample string `validate:"required,oneof=sample hoge"`
}

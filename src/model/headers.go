package model

// HeaderParams ... リクエストヘッダーで受け取る値
type HeaderParams struct {
	Beego string `validate:"required,oneof=beego hoge"`
}

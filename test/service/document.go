package service

import "github.com/aikizoku/beego/test/model"

// Document ... ドキュメントを操作する
type Document interface {
	RemoveAll()
	Distributes(tmplPath string, apis []*model.API)
}

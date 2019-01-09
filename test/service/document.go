package service

import "github.com/aikizoku/gocci/test/model"

// Document ... ドキュメントを操作する
type Document interface {
	RemoveAll()
	Distributes(tmplPath string, apis []*model.API)
}

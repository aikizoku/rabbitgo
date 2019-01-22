package service

import "github.com/aikizoku/skgo/test/model"

// Document ... ドキュメントを操作する
type Document interface {
	RemoveAll()
	Distributes(tmplPath string, apis []*model.API)
}

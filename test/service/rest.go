package service

import "github.com/aikizoku/skgo/test/model"

// Rest ... REST形式のAPI通信を行う
type Rest interface {
	Get(name string, uri string, params map[string]interface{})
	Post(name string, uri string, params map[string]interface{})
	Put(name string, uri string, params map[string]interface{})
	Delete(name string, uri string, params map[string]interface{})
	GetAPIs() []*model.API
}

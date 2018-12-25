package scenario

import "github.com/aikizoku/beego/test/service"

type abnormal struct {
	dSvc service.Document
	rSvc service.Rest
	jSvc service.JSONRPC2
}

func (c *abnormal) Run() {

}

// NewAbnormal ... Abnormalを作成する
func NewAbnormal(
	dSvc service.Document,
	rSvc service.Rest,
	jSvc service.JSONRPC2) Interfaces {
	return &abnormal{
		dSvc: dSvc,
		rSvc: rSvc,
		jSvc: jSvc,
	}
}

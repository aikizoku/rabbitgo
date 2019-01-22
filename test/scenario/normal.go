package scenario

import "github.com/aikizoku/skgo/test/service"

type normal struct {
	dSvc service.Document
	rSvc service.Rest
	jSvc service.JSONRPC2
}

func (s *normal) Run() {
	// リクエスト送る
	s.Send()

	// リクエストとレスポンスをまとめる
	apis := s.rSvc.GetAPIs()

	// お掃除
	s.dSvc.RemoveAll()

	// フォーマットに適用
	s.dSvc.Distributes("template/api.tmpl", apis)
}

func (s *normal) Send() {
	s.rSvc.Post("サンプル", "/sample", map[string]interface{}{
		"sample_id": "12345",
	})
}

// NewNormal ... Normalを作成する
func NewNormal(
	dSvc service.Document,
	rSvc service.Rest,
	jSvc service.JSONRPC2) Interfaces {
	return &normal{
		dSvc: dSvc,
		rSvc: rSvc,
		jSvc: jSvc,
	}
}

package api

import (
	"context"
	"encoding/json"

	"github.com/aikizoku/go-gae-template/src/service"
)

type TestGetHandler struct {
	Svc service.Sample
}

func (s *TestGetHandler) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	return nil, nil
}

func (s *TestGetHandler) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	s.Svc.TestGet(ctx)
	return "success", nil
}

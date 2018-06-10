package api

import (
	"context"
	"encoding/json"

	"github.com/aikizoku/go-gae-template/src/service"
)

type TestPutHandler struct {
	Svc service.Sample
}

func (s *TestPutHandler) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	return nil, nil
}

func (s *TestPutHandler) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	s.Svc.TestPut(ctx)
	return "success", nil
}

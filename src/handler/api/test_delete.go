package api

import (
	"context"
	"encoding/json"

	"github.com/aikizoku/go-gae-template/src/service"
)

type TestDeleteHandler struct {
	Svc service.Sample
}

func (s *TestDeleteHandler) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	return nil, nil
}

func (s *TestDeleteHandler) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	s.Svc.TestDelete(ctx)
	return "success", nil
}

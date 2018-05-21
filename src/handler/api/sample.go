package api

import (
	"context"
	"encoding/json"

	"github.com/aikizoku/go-gae-template/src/model"
	"github.com/aikizoku/go-gae-template/src/service"
)

type SampleHandler struct {
	Service service.Sample
}

func (s *SampleHandler) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	var params model.Sample
	err := json.Unmarshal(*msg, &params)
	return params, err
}

func (s *SampleHandler) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	s.Service.Hoge(ctx)
	sample := params.(model.Sample)
	return []model.Sample{
		sample,
		sample,
	}, nil
}

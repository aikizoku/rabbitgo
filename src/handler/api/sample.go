package api

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aikizoku/go-gae-template/src/model"
	"github.com/aikizoku/go-gae-template/src/service"
	"google.golang.org/appengine/log"
)

type SampleHandler struct {
	Svc service.Sample
}

func (s *SampleHandler) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	var params model.Sample
	err := json.Unmarshal(*msg, &params)
	return params, err
}

func (s *SampleHandler) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	log.Debugf(ctx, "call handler sample")
	s.Svc.Hoge(ctx)
	sample := params.(model.Sample)
	sample.UpdatedAt = time.Now().Unix()
	return []model.Sample{
		sample,
		sample,
	}, nil
}

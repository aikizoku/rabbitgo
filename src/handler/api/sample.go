package api

import (
	"context"
	"encoding/json"
	"time"
)

type Sample struct {
}

func (s *Sample) ParseParams(ctx context.Context, params *json.RawMessage) (interface{}, error) {

}

func (s *Sample) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	a := time.Time()
}

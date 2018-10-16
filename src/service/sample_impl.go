package service

import (
	"context"

	"github.com/aikizoku/beego/src/lib/util"
	"github.com/aikizoku/beego/src/model"
	"github.com/aikizoku/beego/src/repository"
	"google.golang.org/appengine/log"
)

type sample struct {
	repo repository.Sample
}

func (s *sample) Sample(ctx context.Context) (model.Sample, error) {
	log.Debugf(ctx, "call service beego")
	return model.Sample{
		ID:        123,
		Name:      "sample",
		Enabled:   true,
		CreatedAt: util.TimeNow().Unix(),
		UpdatedAt: util.TimeNow().Unix(),
	}, nil
}

func (s *sample) TestDataStore(ctx context.Context) error {
	return nil
}

func (s *sample) TestCloudSQL(ctx context.Context) error {
	return nil
}

func (s *sample) TestHTTP(ctx context.Context) error {
	return nil
}

// NewSample ... サンプルサービスを取得する
func NewSample(repo repository.Sample) Sample {
	return &sample{
		repo: repo,
	}
}

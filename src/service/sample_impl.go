package service

import (
	"context"

	"github.com/aikizoku/beego/src/lib/util"
	"github.com/aikizoku/beego/src/model"
	"github.com/aikizoku/beego/src/repository"
	"google.golang.org/appengine/log"
)

type sampleService struct {
	repo repository.SampleRepository
}

func (s *sampleService) Sample(ctx context.Context) (model.Sample, error) {
	log.Debugf(ctx, "call service beego")
	return model.Sample{
		ID:        123,
		Name:      "sample",
		Enabled:   true,
		CreatedAt: util.TimeNow().Unix(),
		UpdatedAt: util.TimeNow().Unix(),
	}, nil
}

func (s *sampleService) TestDataStore(ctx context.Context) error {
	return nil
}

func (s *sampleService) TestCloudSQL(ctx context.Context) error {
	return nil
}

func (s *sampleService) TestHTTP(ctx context.Context) error {
	return nil
}

// NewSampleService ... サンプルサービスを取得する
func NewSampleService(repo repository.SampleRepository) SampleService {
	return &sampleService{
		repo: repo,
	}
}

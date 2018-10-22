package service

import (
	"context"

	"github.com/aikizoku/beego/src/lib/log"
	"github.com/aikizoku/beego/src/model"
	"github.com/aikizoku/beego/src/repository"
)

type sample struct {
	repo repository.Sample
}

func (s *sample) Sample(ctx context.Context) (model.Sample, error) {
	log.Debugf(ctx, "call service beego")
	return model.Sample{
		ID:       123,
		Category: "hoge",
		Name:     "sample太郎",
		Enabled:  true,
	}, nil
}

func (s *sample) TestDataStore(ctx context.Context) error {
	ids, err := s.repo.DataStoreUpsertMulti(ctx, []model.Sample{
		model.Sample{
			Category: "hoge",
			Name:     "sample太郎",
			Enabled:  true,
		},
		model.Sample{
			Category: "hoge",
			Name:     "sample花子",
			Enabled:  true,
		},
		model.Sample{
			Category: "fuga",
			Name:     "sample佳子",
			Enabled:  true,
		},
	})
	if err != nil {
		return err
	}
	log.Debugf(ctx, "%v", ids)

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

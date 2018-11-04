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
		Category:  "hoge",
		Name:      "sample太郎",
		Enabled:   true,
		CreatedAt: util.TimeNow(),
	}, nil
}

func (s *sample) TestDataStore(ctx context.Context) error {
	ids, err := s.repo.DataStoreUpsertMulti(ctx, []*model.Sample{
		&model.Sample{
			Category:  "hoge",
			Name:      "sample太郎",
			Enabled:   true,
			CreatedAt: util.TimeNow(),
		},
		&model.Sample{
			Category:  "hoge",
			Name:      "sample花子",
			Enabled:   true,
			CreatedAt: util.TimeNow(),
		},
		&model.Sample{
			Category:  "fuga",
			Name:      "sample佳子",
			Enabled:   true,
			CreatedAt: util.TimeNow(),
		},
		&model.Sample{
			Category:  "fuga",
			Name:      "sample忠生",
			Enabled:   false,
			CreatedAt: util.TimeNow(),
		},
	})
	if err != nil {
		return err
	}

	log.Debugf(ctx, "%v", ids)

	dsts, err := s.repo.DataStoreGetByQuery(ctx, "fuga")
	if err != nil {
		return err
	}
	for _, dst := range dsts {
		log.Debugf(ctx, "%v", dst)
	}

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

package service

import (
	"context"

	"github.com/aikizoku/beego/src/model"
	"github.com/aikizoku/beego/src/repository"
	"github.com/aikizoku/beego/src/util"
	"google.golang.org/appengine/log"
)

type beego struct {
	repo repository.Beego
}

func (s *beego) Beego(ctx context.Context) (model.Beego, error) {
	log.Debugf(ctx, "call service beego")
	return model.Beego{
		ID:   123,
		Name: "beego",
		Type: model.BeegoType{
			ID:        456,
			Name:      "beego_type",
			CreatedAt: util.TimeNow(),
			UpdatedAt: util.TimeNow(),
		},
		Enabled:   true,
		CreatedAt: util.TimeNow(),
		UpdatedAt: util.TimeNow(),
	}
}

// NewBeego ... サービスを取得する
func NewBeego(repo repository.Beego) Beego {
	return &beego{
		repo: repo,
	}
}

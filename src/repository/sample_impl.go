package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/aikizoku/go-gae-template/src/model"
	"google.golang.org/appengine/datastore"

	"github.com/mjibson/goon"

	"github.com/aikizoku/go-gae-template/src/infrastructure"
	"google.golang.org/appengine/log"
)

const (
	TYPE_NEW       = "new"
	coutf8LastChar = "\xef\xbf\xbd"
)

type sample struct {
	http *infrastructure.HTTP
	// csql *sql.DB
}

func (s *sample) Hoge(ctx context.Context) {
	log.Debugf(ctx, "call repository hoge")
}

func (s *sample) TestPut(ctx context.Context) {
	now := time.Now().Unix()
	client := goon.FromContext(ctx)

	vs := []model.ArticleList{}
	for i := 0; i < 19; i++ {
		v := model.ArticleList{
			ID:          123,
			Type:        TYPE_NEW,
			Title:       fmt.Sprintf("title_%d", i),
			Description: fmt.Sprintf("description_%d", i),
			PublishedAt: now,
		}
		vs = append(vs, v)
	}

	keys, err := client.PutMulti(vs)
	if err != nil {
		log.Errorf(ctx, err.Error())
		return
	}
	log.Infof(ctx, "%v", keys)

	// err = client.Get(&model.Sample{ID: 123})
	// if err != nil {
	// 	log.Errorf(ctx, err.Error())
	// 	return
	// }
	// log.Infof(ctx, "%v", key)
}

func (s *sample) TestGet(ctx context.Context) {
	client := goon.FromContext(ctx)
	q := datastore.NewQuery("ArticleList").Filter("type >=", "new").Filter("type <=", "new"+coutf8LastChar).Order("type").Limit(40)
	var r []*model.ArticleList
	_, err := client.GetAll(q, &r)
	// t := client.Run(q)
	if err != nil {
		log.Errorf(ctx, err.Error())
		return
	}

	for _, ret := range r {
		log.Infof(ctx, "%s", ret.Title)
	}
}

func (s *sample) TestDelete(ctx context.Context) {
	// client := goon.FromContext(ctx)
}

// NewSample ...
func NewSample(http *infrastructure.HTTP) Sample {
	return &sample{
		http: http,
		// csql: csql,
	}
}

package bigquery

import (
	"context"
	"reflect"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"github.com/aikizoku/rabbitgo/appengine/src/lib/log"
)

// Client ... BigQueryのクライアント
type Client struct {
	client *bigquery.Client
}

// GetList ... クエリを実行し、データを取得する
func (c *Client) GetList(ctx context.Context, query string, limit int, cursor string, dsts interface{}) (string, error) {
	q := c.client.Query(query)
	it, err := q.Read(ctx)
	if err != nil {
		log.Errorm(ctx, "q.Read", err)
		return "", err
	}
	it.PageInfo().MaxSize = limit
	it.PageInfo().Token = cursor

	rv := reflect.Indirect(reflect.ValueOf(dsts))
	rrt := rv.Type().Elem().Elem()
	i := 0
	for {
		i++
		v := reflect.New(rrt).Interface()
		err = it.Next(v)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Errorm(ctx, "it.Next", err)
			return "", err
		}
		rrv := reflect.ValueOf(v)
		rv.Set(reflect.Append(rv, rrv))
		if i == limit {
			break
		}
	}
	return it.PageInfo().Token, nil
}

// NewClient ... クライアントを作成する
func NewClient(projectID string, credentialsPath string) *Client {
	ctx := context.Background()
	opt := option.WithCredentialsFile(credentialsPath)
	client, err := bigquery.NewClient(ctx, projectID, opt)
	if err != nil {
		panic(err)
	}
	return &Client{
		client: client,
	}
}

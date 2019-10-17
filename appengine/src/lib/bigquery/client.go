package bigquery

import (
	"context"
	"reflect"
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

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
	if pageInfo := it.PageInfo(); pageInfo != nil {
		pageInfo.MaxSize = limit
		pageInfo.Token = cursor
	}

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
	var token string
	if pageInfo := it.PageInfo(); pageInfo != nil {
		token = pageInfo.Token
	}
	return token, nil
}

// NewClient ... クライアントを作成する
func NewClient(projectID string, credentialsPath string) *Client {
	ctx := context.Background()
	cOpt := option.WithCredentialsFile(credentialsPath)
	gOpt := option.WithGRPCDialOption(grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                30 * time.Millisecond,
		Timeout:             20 * time.Millisecond,
		PermitWithoutStream: true,
	}))
	client, err := bigquery.NewClient(ctx, projectID, cOpt, gOpt)
	if err != nil {
		panic(err)
	}
	return &Client{
		client: client,
	}
}

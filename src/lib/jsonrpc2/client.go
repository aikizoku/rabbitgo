package jsonrpc2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aikizoku/skgo/src/lib/httpclient"
	"github.com/aikizoku/skgo/src/lib/log"
)

// Client ... JSONRPC2のリクエストを行う
type Client struct {
	URL      string
	Headers  map[string]string
	Requests []*ClientRequest
}

// AddRequest ... JSONRPC2のリクエストを登録する
func (c *Client) AddRequest(ctx context.Context, id string, method string, params interface{}) error {
	rawParams, err := c.marshalRawMessage(ctx, params)
	if err != nil {
		log.Errorm(ctx, "c.marshalRawMessage", err)
		return err
	}
	c.Requests = append(c.Requests, &ClientRequest{
		Version: version,
		ID:      id,
		Method:  method,
		Params:  rawParams,
	})
	return nil
}

// DoSingle ... JSONRPC2のシングルリクエストを行う
func (c *Client) DoSingle(ctx context.Context, method string, params interface{}) (*json.RawMessage, *ErrorResponse, error) {
	rawParams, err := c.marshalRawMessage(ctx, params)
	if err != nil {
		log.Errorm(ctx, "c.marshalRawMessage", err)
		return nil, nil, err
	}

	req := &ClientRequest{
		Version: version,
		ID:      "single",
		Method:  method,
		Params:  rawParams,
	}
	rawReq, err := json.Marshal(req)
	if err != nil {
		log.Errorm(ctx, "json.Marshal", err)
		return nil, nil, err
	}

	status, body, err := httpclient.PostJSON(ctx, c.URL, rawReq, &httpclient.HTTPOption{Headers: c.Headers})
	if err != nil {
		log.Errorm(ctx, "httpclient.PostJSON", err)
		return nil, nil, err
	}
	if status != http.StatusOK {
		err = fmt.Errorf("httpclient.PostJSON status: %d", status)
		log.Errorf(ctx, err.Error())
		return nil, nil, err
	}

	var res ClientResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Errorm(ctx, "json.Unmarshal", err)
		return nil, nil, err
	}

	return res.Result, res.Error, nil
}

// DoBatch ... JSONRPC2のバッチリクエストを行う
func (c *Client) DoBatch(ctx context.Context) ([]*ClientResponse, error) {
	p, err := json.Marshal(c.Requests)
	if err != nil {
		log.Errorm(ctx, "json.Marshal", err)
		return nil, err
	}

	status, body, err := httpclient.PostJSON(ctx, c.URL, p, nil)
	if err != nil {
		log.Errorm(ctx, "httpclient.PostJSON", err)
		return nil, err
	}
	if status != http.StatusOK {
		err = fmt.Errorf("httpclient.PostJSON status: %d", status)
		log.Errorf(ctx, err.Error())
		return nil, err
	}

	var res []*ClientResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Errorm(ctx, "json.Unmarshal", err)
		return nil, err
	}

	return res, nil
}

func (c *Client) marshalRawMessage(ctx context.Context, params interface{}) (*json.RawMessage, error) {
	bParams, err := json.Marshal(params)
	if err != nil {
		log.Errorm(ctx, "json.Marshal", err)
		return nil, err
	}
	rawMessage := json.RawMessage(bParams)
	return &rawMessage, nil
}

// NewClient ... Clientを作成する
func NewClient(url string, headers map[string]string) *Client {
	return &Client{
		URL:     url,
		Headers: headers,
	}
}

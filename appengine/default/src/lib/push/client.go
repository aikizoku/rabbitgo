package push

import (
	"context"

	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/jsonrpc2"
	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/log"
)

type entryParams struct {
	AppID    string   `json:"app_id"`
	UserID   string   `json:"user_id"`
	Platform Platform `json:"platform"`
	DeviceID string   `json:"device_id"`
	Token    string   `json:"token"`
}

type leaveParams struct {
	AppID    string   `json:"app_id"`
	UserID   string   `json:"user_id"`
	Platform Platform `json:"platform"`
	DeviceID string   `json:"device_id"`
}

type sendParams struct {
	AppID   string   `json:"app_id"`
	UserIDs []string `json:"user_ids"`
	Message *Message `json:"message"`
}

// Client ... プッシュ通知のクライアント
type Client struct {
	appID    string
	endpoint string
	headers  map[string]string
}

// Entry ... 登録する
func (c *Client) Entry(ctx context.Context, userID string, pf Platform, deviceID string, token string) error {
	params := &entryParams{
		AppID:    c.appID,
		UserID:   userID,
		Platform: pf,
		DeviceID: deviceID,
		Token:    token,
	}
	client := jsonrpc2.NewClient(c.endpoint, c.headers)
	_, resError, err := client.DoSingle(ctx, "entry", params)
	if err != nil {
		log.Errorm(ctx, "client.DoSingle", err)
		return err
	}
	if resError != nil {
		err = log.Errore(ctx, "code: %d, message: %s", resError.Code, resError.Message)
		return err
	}
	return nil
}

// Leave ... 解除する
func (c *Client) Leave(ctx context.Context, userID string, pf Platform, deviceID string) error {
	params := &leaveParams{
		AppID:    c.appID,
		UserID:   userID,
		Platform: pf,
		DeviceID: deviceID,
	}
	client := jsonrpc2.NewClient(c.endpoint, c.headers)
	_, resError, err := client.DoSingle(ctx, "leave", params)
	if err != nil {
		log.Errorm(ctx, "client.DoSingle", err)
		return err
	}
	if resError != nil {
		err = log.Errore(ctx, "code: %d, message: %s", resError.Code, resError.Message)
		return err
	}
	return nil
}

// Send ... 送信する
func (c *Client) Send(ctx context.Context, userIDs []string, msg *Message) error {
	params := &sendParams{
		AppID:   c.appID,
		UserIDs: userIDs,
		Message: msg,
	}
	cli := jsonrpc2.NewClient(c.endpoint, c.headers)
	_, resError, err := cli.DoSingle(ctx, "send", params)
	if err != nil {
		log.Errorm(ctx, "cli.DoSingle", err)
		return err
	}
	if resError != nil {
		err = log.Errore(ctx, "code: %d, message: %s", resError.Code, resError.Message)
		return err
	}
	return nil
}

// NewClient ... クライアントを作成する
func NewClient(appID string, endpoint string, authToken string) *Client {
	headers := map[string]string{
		"Authorization": authToken,
	}
	return &Client{
		appID:    appID,
		endpoint: endpoint,
		headers:  headers,
	}
}

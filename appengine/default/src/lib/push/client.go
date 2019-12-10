package push

import (
	"context"
	"encoding/json"

	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/jsonrpc2"
	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/log"
)

// Client ... プッシュ通知のクライアント
type Client struct {
	appID    string
	endpoint string
	headers  map[string]string
}

// Entry ... 登録する
func (c *Client) Entry(ctx context.Context, userID string, pf Platform, deviceID string, token string) error {
	params := &struct {
		AppID    string   `json:"app_id"`
		UserID   string   `json:"user_id"`
		Platform Platform `json:"platform"`
		DeviceID string   `json:"device_id"`
		Token    string   `json:"token"`
	}{
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
	params := &struct {
		AppID    string   `json:"app_id"`
		UserID   string   `json:"user_id"`
		Platform Platform `json:"platform"`
		DeviceID string   `json:"device_id"`
	}{
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

// SendByUsers ... 指定したユーザーに送信する
func (c *Client) SendByUsers(ctx context.Context, userIDs []string, msg *Message) error {
	params := &struct {
		AppID   string   `json:"app_id"`
		UserIDs []string `json:"user_ids"`
		Message *Message `json:"message"`
	}{
		AppID:   c.appID,
		UserIDs: userIDs,
		Message: msg,
	}
	cli := jsonrpc2.NewClient(c.endpoint, c.headers)
	_, resError, err := cli.DoSingle(ctx, "send_by_users", params)
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

// SendByAllUsers ... 全員に送信する
func (c *Client) SendByAllUsers(ctx context.Context, msg *Message) error {
	params := &struct {
		AppID   string   `json:"app_id"`
		Message *Message `json:"message"`
	}{
		AppID:   c.appID,
		Message: msg,
	}
	cli := jsonrpc2.NewClient(c.endpoint, c.headers)
	_, resError, err := cli.DoSingle(ctx, "send_by_all_users", params)
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

// GetReserve ... 予約を取得する
func (c *Client) GetReserve(ctx context.Context, reserveID string) (*Reserve, error) {
	params := &struct {
		AppID     string `json:"app_id"`
		ReserveID string `json:"reserve_id"`
	}{
		AppID:     c.appID,
		ReserveID: reserveID,
	}
	cli := jsonrpc2.NewClient(c.endpoint, c.headers)
	resResult, resError, err := cli.DoSingle(ctx, "get_reserve", params)
	if err != nil {
		log.Errorm(ctx, "cli.DoSingle", err)
		return nil, err
	}
	if resError != nil {
		err = log.Errore(ctx, "code: %d, message: %s", resError.Code, resError.Message)
		return nil, err
	}
	ret := &struct {
		Reserve *Reserve `json:"reserve"`
	}{}
	err = json.Unmarshal([]byte(*resResult), ret)
	if err != nil {
		log.Errorm(ctx, "json.Unmarshal", err)
		return nil, err
	}
	return ret.Reserve, nil
}

// ListReserve ... 予約リストを取得する
func (c *Client) ListReserve(ctx context.Context, limit int, cursor string) ([]*Reserve, string, error) {
	params := &struct {
		AppID  string `json:"app_id"`
		Limit  int    `json:"limit"`
		Cursor string `json:"cursor"`
	}{
		AppID:  c.appID,
		Limit:  limit,
		Cursor: cursor,
	}
	cli := jsonrpc2.NewClient(c.endpoint, c.headers)
	resResult, resError, err := cli.DoSingle(ctx, "list_reserve", params)
	if err != nil {
		log.Errorm(ctx, "cli.DoSingle", err)
		return nil, "", err
	}
	if resError != nil {
		err = log.Errore(ctx, "code: %d, message: %s", resError.Code, resError.Message)
		return nil, "", err
	}
	ret := &struct {
		Reserves   []*Reserve `json:"reserves"`
		NextCursor string     `json:"next_cursor"`
	}{}
	err = json.Unmarshal([]byte(*resResult), ret)
	if err != nil {
		log.Errorm(ctx, "json.Unmarshal", err)
		return nil, "", err
	}
	return ret.Reserves, ret.NextCursor, nil
}

// CreateReserve ... 予約リストを作成する
func (c *Client) CreateReserve(ctx context.Context, msg *Message, reservedAt int64) (*Reserve, error) {
	params := &struct {
		AppID      string   `json:"app_id"`
		Message    *Message `json:"message"`
		ReservedAt int64    `json:"reserved_at"`
	}{
		AppID:      c.appID,
		Message:    msg,
		ReservedAt: reservedAt,
	}
	cli := jsonrpc2.NewClient(c.endpoint, c.headers)
	resResult, resError, err := cli.DoSingle(ctx, "create_reserve", params)
	if err != nil {
		log.Errorm(ctx, "cli.DoSingle", err)
		return nil, err
	}
	if resError != nil {
		err = log.Errore(ctx, "code: %d, message: %s", resError.Code, resError.Message)
		return nil, err
	}
	ret := &struct {
		Reserve *Reserve `json:"reserve"`
	}{}
	err = json.Unmarshal([]byte(*resResult), ret)
	if err != nil {
		log.Errorm(ctx, "json.Unmarshal", err)
		return nil, err
	}
	return ret.Reserve, nil
}

// UpdateReserve ... 予約リストを更新する
func (c *Client) UpdateReserve(ctx context.Context, reserveID string, msg *Message, reservedAt int64, status ReserveStatus) (*Reserve, error) {
	params := &struct {
		AppID      string        `json:"app_id"`
		ReserveID  string        `json:"reserve_id"`
		Message    *Message      `json:"message"`
		ReservedAt int64         `json:"reserved_at"`
		Status     ReserveStatus `json:"status"`
	}{
		AppID:      c.appID,
		ReserveID:  reserveID,
		Message:    msg,
		ReservedAt: reservedAt,
		Status:     status,
	}
	cli := jsonrpc2.NewClient(c.endpoint, c.headers)
	resResult, resError, err := cli.DoSingle(ctx, "update_reserve", params)
	if err != nil {
		log.Errorm(ctx, "cli.DoSingle", err)
		return nil, err
	}
	if resError != nil {
		err = log.Errore(ctx, "code: %d, message: %s", resError.Code, resError.Message)
		return nil, err
	}
	ret := &struct {
		Reserve *Reserve `json:"reserve"`
	}{}
	err = json.Unmarshal([]byte(*resResult), ret)
	if err != nil {
		log.Errorm(ctx, "json.Unmarshal", err)
		return nil, err
	}
	return ret.Reserve, nil
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

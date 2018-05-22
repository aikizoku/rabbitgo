package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine/log"
)

const (
	// Jsonrpc2ErrInvalidRequest ... 不正なリクエスト
	Jsonrpc2ErrInvalidRequest = 40001
	// Jsonrpc2ErrInvalidJSON ... 不正なJSON形式
	Jsonrpc2ErrInvalidJSON = 40002
	// Jsonrpc2ErrInvalidJsonrpc2 ... JSONがJSONRPC2の形式ではない
	Jsonrpc2ErrInvalidJsonrpc2 = 40003
	// Jsonrpc2ErrInvalidParams ... 不正なパラメータ
	Jsonrpc2ErrInvalidParams = 40004
	// Jsonrpc2ErrMehodNotFaund ... 存在しないMethod
	Jsonrpc2ErrMehodNotFaund = 40401
	// Jsonrpc2ErrInternal ... 内部エラー
	Jsonrpc2ErrInternal = 50001

	contentType = "application/json"
	version     = "2.0"
)

// Jsonrpc2Handler ... JSORPC2ハンドラの定義
type Jsonrpc2Handler interface {
	DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error)
	Exec(ctx context.Context, method string, params interface{}) (interface{}, error)
}

// Jsonrpc2 ... JSONRPC2に準拠したライブラリ
type Jsonrpc2 struct {
	handlers map[string]Jsonrpc2Handler
}

// NewJsonrpc2 ... JSONRPC2を作成する
func NewJsonrpc2() *Jsonrpc2 {
	return &Jsonrpc2{
		handlers: map[string]Jsonrpc2Handler{},
	}
}

// Register ... JSONRPC2のリクエストを登録する
func (j *Jsonrpc2) Register(method string, handler Jsonrpc2Handler) {
	if method == "" || handler == nil {
		return
	}
	j.handlers[method] = handler
}

// Handle ... JSONRPC2のリクエストをハンドルする
func (j *Jsonrpc2) Handle(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)

	// POSTで送信されていること
	if r.Method != "POST" {
		log.Errorf(ctx, "invalid http method: %s", r.Method)
		RenderJSON(w, http.StatusNotAcceptable, nil)
		return
	}

	// リクエストのContent-TypeもしくはAcceptがapplication/jsonであること
	contentType := r.Header.Get("Content-Type")
	if contentType != contentType {
		log.Errorf(ctx, "invalid http header content-type: %s", contentType)
		RenderJSON(w, http.StatusUnsupportedMediaType, nil)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf(ctx, "failed to read http body")
		RenderJSON(w, http.StatusBadRequest, nil)
		return
	}

	err = j.handleSingleRequest(ctx, w, r, data)
	if err != nil {
		err = j.handleBatchRequest(ctx, w, r, data)
	}
	if err != nil {
		log.Errorf(ctx, "failed to parse json")
		RenderJSON(w, http.StatusBadRequest, nil)
		return
	}
}

func (j *Jsonrpc2) handleSingleRequest(ctx context.Context, w http.ResponseWriter, r *http.Request, data []byte) error {
	var req jsonrpc2Request
	err := json.Unmarshal(data, &req)
	if err != nil {
		return err
	}
	res := j.handleRequest(ctx, r, req)
	encoder := json.NewEncoder(w)
	return encoder.Encode(res)
}

func (j *Jsonrpc2) handleBatchRequest(ctx context.Context, w http.ResponseWriter, r *http.Request, data []byte) error {
	var reqs []jsonrpc2Request
	err := json.Unmarshal(data, &reqs)
	if err != nil {
		return err
	}
	var responses []jsonrpc2Response
	ch := make(chan jsonrpc2Response, len(reqs))
	for _, req := range reqs {
		go func(_req jsonrpc2Request) {
			ch <- j.handleRequest(ctx, r, _req)
		}(req)
	}
	for i := 0; i < len(reqs); i++ {
		responses = append(responses, <-ch)
	}
	encoder := json.NewEncoder(w)
	return encoder.Encode(responses)
}

func (j *Jsonrpc2) handleRequest(ctx context.Context, r *http.Request, req jsonrpc2Request) jsonrpc2Response {
	if !req.isValid() {
		msg := fmt.Sprintf("invalid jsonrpc2 params: %v", req)
		log.Errorf(ctx, msg)
		return newJsonrpc2ErrorResponse(req.ID, Jsonrpc2ErrInvalidJsonrpc2, msg)
	}
	handler := j.handlers[req.Method]
	if handler == nil {
		msg := fmt.Sprintf("method not found: %s", req.Method)
		log.Errorf(ctx, msg)
		return newJsonrpc2ErrorResponse(req.ID, Jsonrpc2ErrMehodNotFaund, msg)
	}
	params, err := handler.DecodeParams(ctx, req.Params)
	if err != nil {
		msg := fmt.Sprintf("invalid params: %v", err)
		log.Errorf(ctx, msg)
		return newJsonrpc2ErrorResponse(req.ID, Jsonrpc2ErrInvalidParams, msg)
	}
	result, err := handler.Exec(ctx, req.Method, params)
	if err != nil {
		msg := fmt.Sprintf("invalid params: %v", err)
		log.Errorf(ctx, msg)
		return newJsonrpc2ErrorResponse(req.ID, Jsonrpc2ErrInternal, msg)
	}
	return newJsonrpc2Response(req.ID, result)
}

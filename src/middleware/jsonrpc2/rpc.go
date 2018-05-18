package jsonrpc2

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aikizoku/go-gae-template/src/handler"
)

const (
	// ErrInvalidRequest ... 不正なリクエスト
	ErrInvalidRequest = 40001
	// ErrInvalidJSON ... 不正なJSON形式
	ErrInvalidJSON = 40002
	// ErrInvalidJsonrpc2 ... JSONがJSONRPC2の形式ではない
	ErrInvalidJsonrpc2 = 40003
	// ErrInvalidParams ... 不正なパラメータ
	ErrInvalidParams = 40004
	// ErrMehodNotFaund ... 存在しないMethod
	ErrMehodNotFaund = 40401
	// ErrInternal ... 内部エラー
	ErrInternal = 50001

	contentType = "application/json"
)

type Handler interface {
	ParseParams(ctx context.Context, params *json.RawMessage) (interface{}, error)
	Exec(ctx context.Context, method string, params interface{}) (interface{}, error)
}

type Jsonrpc2 struct {
	handlers map[string]Handler
}

func NewJsonrpc2() *Jsonrpc2 {
	return &Jsonrpc2{
		handlers: map[string]Handler{},
	}
}

func (j *Jsonrpc2) Register(method string, handler Handler) {
	if method == "" || handler == nil {
		return
	}
	j.handlers[method] = handler
}

func (j *Jsonrpc2) Handle(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)

	// POSTで送信されていること
	if r.Method != "POST" {
		msg := fmt.Sprintf("invalid http method: %s", r.Method)
		handler.RenderJSON(w, http.StatusNotAcceptable, newErrorResponse("", ErrInvalidRequest, msg))
		return
	}

	// リクエストのContent-TypeもしくはAcceptがapplication/jsonであること
	contentType := r.Header.Get("Content-Type")
	accept := r.Header.Get("Accept")
	if contentType != contentType || accept != contentType {
		msg := fmt.Sprintf("invalid http header content-type: %s, accept: %s", contentType, accept)
		handler.RenderJSON(w, http.StatusUnsupportedMediaType, newErrorResponse("", ErrInvalidRequest, msg))
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "failed to read http body"
		handler.RenderJSON(w, http.StatusBadRequest, newErrorResponse("", ErrInvalidRequest, msg))
		return
	}

	err = j.handleSingleRequest(ctx, w, r, data)
	if err != nil {
		err = j.handleBatchRequest(ctx, w, r, data)
	}
	if err != nil {
		msg := "failed to parse json"
		handler.RenderJSON(w, http.StatusBadRequest, newErrorResponse("", ErrInvalidJSON, msg))
		return
	}
}

func (j *Jsonrpc2) handleSingleRequest(ctx context.Context, w http.ResponseWriter, r *http.Request, data []byte) error {
	var req request
	err := json.Unmarshal(data, &req)
	if err != nil {
		return err
	}
	res := j.handleRequest(ctx, r, req)
	encoder := json.NewEncoder(w)
	return encoder.Encode(res)
}

func (j *Jsonrpc2) handleBatchRequest(ctx context.Context, w http.ResponseWriter, r *http.Request, data []byte) error {
	var reqs []request
	err := json.Unmarshal(data, &reqs)
	if err != nil {
		return err
	}
	var responses []response
	ch := make(chan response, len(reqs))
	for _, req := range reqs {
		go func(_req request) {
			ch <- j.handleRequest(ctx, r, _req)
		}(req)
	}
	for i := 0; i < len(reqs); i++ {
		responses = append(responses, <-ch)
	}
	encoder := json.NewEncoder(w)
	return encoder.Encode(responses)
}

func (j *Jsonrpc2) handleRequest(ctx context.Context, r *http.Request, req request) response {
	if !req.isValid() {
		msg := "invalid jsonrpc2 params"
		return newErrorResponse(req.ID, ErrInvalidJsonrpc2, msg)
	}
	handler := j.handlers[req.Method]
	if handler == nil {
		msg := fmt.Sprintf("method not found: %s", req.Method)
		return newErrorResponse(req.ID, ErrMehodNotFaund, msg)
	}
	params, err := handler.ParseParams(ctx, req.Params)
	if err != nil {
		msg := fmt.Sprintf("invalid params: %v", err)
		return newErrorResponse(req.ID, ErrInvalidParams, msg)
	}
	result, err := handler.Exec(ctx, req.Method, params)
	if err != nil {
		msg := fmt.Sprintf("invalid params: %v", err)
		return newErrorResponse(req.ID, ErrInternal, msg)
	}
	return newResponse(req.ID, result)
}

package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"google.golang.org/appengine"
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

	jsonrpc2ContentType = "application/json"
	jsonrpc2Version     = "2.0"
)

// Jsonrpc2 ... JSONRPC2に準拠したライブラリ
type Jsonrpc2 struct {
	handlers map[string]Jsonrpc2Handler
}

// Jsonrpc2Handler ... JSORPC2ハンドラの定義
type Jsonrpc2Handler interface {
	DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error)
	Exec(ctx context.Context, method string, params interface{}) (interface{}, error)
}

// Register ... JSONRPC2のリクエストを登録する
func (j *Jsonrpc2) Register(method string, handler Jsonrpc2Handler) {
	if method == "" || handler == nil {
		return
	}
	j.handlers[method] = handler
}

// Handle ... JSONRPC2のリクエストをハンドルする
func (j *Jsonrpc2) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		w.Header().Set("Content-Type", jsonrpc2ContentType)

		// POSTで送信されていること
		if r.Method != "POST" {
			j.renderError(ctx, w, http.StatusNotAcceptable, "invalid http method: %s", r.Method)
			return
		}

		// リクエストのContent-TypeもしくはAcceptがapplication/jsonであること
		contentType := r.Header.Get("Content-Type")
		if contentType != contentType {
			j.renderError(ctx, w, http.StatusUnsupportedMediaType, "invalid http header content-type: %s", contentType)
			return
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			j.renderError(ctx, w, http.StatusUnsupportedMediaType, "read http body error: %s", err.Error())
			return
		}

		err = j.handleSingleRequest(ctx, w, r, data)
		if err != nil {
			err = j.handleBatchRequest(ctx, w, r, data)
		}
		if err != nil {
			j.renderError(ctx, w, http.StatusBadRequest, "parse json error: %s", err.Error())
			return
		}
		next.ServeHTTP(w, r)
	})
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
		return j.renderErrorJSON(ctx, req.ID, Jsonrpc2ErrInvalidJsonrpc2, "invalid jsonrpc2 params: %s", spew.Sdump(req))
	}
	handler := j.handlers[req.Method]
	if handler == nil {
		return j.renderErrorJSON(ctx, req.ID, Jsonrpc2ErrMehodNotFaund, "method not found: %s", req.Method)
	}
	params, err := handler.DecodeParams(ctx, req.Params)
	if err != nil {
		return j.renderErrorJSON(ctx, req.ID, Jsonrpc2ErrInvalidParams, "invalid params: %s", err.Error())
	}
	result, err := handler.Exec(ctx, req.Method, params)
	if err != nil {
		return j.renderErrorJSON(ctx, req.ID, Jsonrpc2ErrInternal, "internal error: %s", err.Error())
	}
	return newJsonrpc2Response(req.ID, result)
}

func (j *Jsonrpc2) renderError(ctx context.Context, w http.ResponseWriter, status int, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a)
	log.Errorf(ctx, msg)
	RenderError(w, status, msg)
}

func (j *Jsonrpc2) renderErrorJSON(ctx context.Context, rpcID string, rpcStatus int, format string, a ...interface{}) jsonrpc2Response {
	msg := fmt.Sprintf(format, a)
	log.Errorf(ctx, msg)
	return newJsonrpc2ErrorResponse(rpcID, rpcStatus, msg)
}

// NewJsonrpc2 ... JSONRPC2を作成する
func NewJsonrpc2() *Jsonrpc2 {
	return &Jsonrpc2{
		handlers: map[string]Jsonrpc2Handler{},
	}
}

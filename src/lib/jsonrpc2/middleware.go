package jsonrpc2

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aikizoku/gocci/src/lib/log"
	"github.com/davecgh/go-spew/spew"
	"github.com/unrolled/render"
)

// Middleware ... JSONRPC2に準拠したミドルウェア
type Middleware struct {
	handlers map[string]Handler
}

// Register ... JSONRPC2のリクエストを登録する
func (m *Middleware) Register(method string, handler Handler) {
	if method == "" || handler == nil {
		return
	}
	m.handlers[method] = handler
}

// Handle ... JSONRPC2のリクエストをハンドルする
func (m *Middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		w.Header().Set("Content-Type", contentType)

		// POSTで送信されていること
		if r.Method != "POST" {
			m.renderError(ctx, w, http.StatusNotAcceptable, "invalid http method: %s", r.Method)
			return
		}

		// リクエストのContent-TypeもしくはAcceptがapplication/jsonであること
		contentType := r.Header.Get("Content-Type")
		if contentType != contentType {
			m.renderError(ctx, w, http.StatusUnsupportedMediaType, "invalid http header content-type: %s", contentType)
			return
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			m.renderError(ctx, w, http.StatusUnsupportedMediaType, "read http body error: %s", err.Error())
			return
		}

		err = m.handleSingleRequest(ctx, w, r, data)
		if err != nil {
			err = m.handleBatchRequest(ctx, w, r, data)
		}
		if err != nil {
			m.renderError(ctx, w, http.StatusBadRequest, "parse json error: %s", err.Error())
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) handleSingleRequest(ctx context.Context, w http.ResponseWriter, r *http.Request, data []byte) error {
	var req request
	err := json.Unmarshal(data, &req)
	if err != nil {
		return err
	}
	res := m.handleRequest(ctx, r, req)
	encoder := json.NewEncoder(w)
	return encoder.Encode(res)
}

func (m *Middleware) handleBatchRequest(ctx context.Context, w http.ResponseWriter, r *http.Request, data []byte) error {
	var reqs []request
	err := json.Unmarshal(data, &reqs)
	if err != nil {
		return err
	}

	var responses []response
	ch := make(chan response, len(reqs))
	for _, req := range reqs {
		go func(_req request) {
			ch <- m.handleRequest(ctx, r, _req)
		}(req)
	}
	for i := 0; i < len(reqs); i++ {
		responses = append(responses, <-ch)
	}

	encoder := json.NewEncoder(w)
	return encoder.Encode(responses)
}

func (m *Middleware) handleRequest(ctx context.Context, r *http.Request, req request) response {
	if !req.isValid() {
		return m.renderErrorJSON(ctx, req.ID, ErrInvalidJsonrpc2, "invalid jsonrpc2 params: %s", spew.Sdump(req))
	}

	handler := m.handlers[req.Method]
	if handler == nil {
		return m.renderErrorJSON(ctx, req.ID, ErrMehodNotFaund, "method not found: %s", req.Method)
	}

	params, err := handler.DecodeParams(ctx, req.Params)
	if err != nil {
		return m.renderErrorJSON(ctx, req.ID, ErrInvalidParams, "invalid params: %s", err.Error())
	}

	result, err := handler.Exec(ctx, req.Method, params)
	if err != nil {
		return m.renderErrorJSON(ctx, req.ID, ErrInternal, "internal error: %s", err.Error())
	}

	return newResponse(req.ID, result)
}

func (m *Middleware) renderError(ctx context.Context, w http.ResponseWriter, status int, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Errorf(ctx, msg)
	render.New().Text(w, status, fmt.Sprintf("%d %s", status, msg))
}

func (m *Middleware) renderErrorJSON(ctx context.Context, rpcID string, rpcStatus int, format string, a ...interface{}) response {
	msg := fmt.Sprintf(format, a...)
	log.Errorf(ctx, msg)
	return newErrorResponse(rpcID, rpcStatus, msg)
}

// NewMiddleware ... Middlewareを作成する
func NewMiddleware() *Middleware {
	return &Middleware{
		handlers: map[string]Handler{},
	}
}

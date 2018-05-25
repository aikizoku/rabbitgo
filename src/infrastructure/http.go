package infrastructure

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

// HTTP ... HTTP通信モジュール
type HTTP struct {
	Timeout time.Duration
}

// HTTPOption ...
type HTTPOption struct {
	Headers map[string]string
	Timeout time.Duration
}

// NewHTTP ... HTTP通信モジュールを作成する
func NewHTTP(timeout time.Duration) HTTP {
	return HTTP{
		Timeout: timeout,
	}
}

// Get ... Getリクエスト(URL)
func (h *HTTP) Get(ctx context.Context, u string, opt *HTTPOption) (bool, int, []byte) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		log.Warningf(ctx, "create request error: %s", err.Error())
		return false, 0, nil
	}
	for key, value := range opt.Headers {
		req.Header.Set(key, value)
	}
	return h.send(ctx, req, opt.Timeout)
}

// GetForm ... Getリクエスト(URL, Params)
func (h *HTTP) GetForm(ctx context.Context, u string, params map[string]string, opt *HTTPOption) (bool, int, []byte) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		log.Warningf(ctx, "create request error: %s", err.Error())
		return false, 0, nil
	}
	for key, value := range opt.Headers {
		req.Header.Set(key, value)
	}
	query := req.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()
	return h.send(ctx, req, opt.Timeout)
}

// GetQueryString ... Getリクエスト(URL, QueryString)
func (h *HTTP) GetQueryString(ctx context.Context, u string, qs string, opt *HTTPOption) (bool, int, []byte) {
	req, err := http.NewRequest("GET", u+"?"+qs, nil)
	if err != nil {
		log.Warningf(ctx, "create request error: %s", err.Error())
		return false, 0, nil
	}
	for key, value := range opt.Headers {
		req.Header.Set(key, value)
	}
	return h.send(ctx, req, opt.Timeout)
}

// PostForm ... Postリクエスト(URL, Params)
func (h *HTTP) PostForm(ctx context.Context, u string, params map[string]string, opt *HTTPOption) (bool, int, []byte) {
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}
	req, err := http.NewRequest("POST", u, strings.NewReader(values.Encode()))
	if err != nil {
		log.Warningf(ctx, "create request error: %s", err.Error())
		return false, 0, nil
	}
	for key, value := range opt.Headers {
		req.Header.Set(key, value)
	}
	return h.send(ctx, req, opt.Timeout)
}

// PostJSON ... Postリクエスト(URL, JSON)
func (h *HTTP) PostJSON(ctx context.Context, u string, json []byte, opt *HTTPOption) (bool, int, []byte) {
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(json))
	if err != nil {
		log.Warningf(ctx, "create request error: %s", err.Error())
		return false, 0, nil
	}
	for key, value := range opt.Headers {
		req.Header.Set(key, value)
	}
	opt.Headers["Content-Type"] = "application/json"
	return h.send(ctx, req, opt.Timeout)
}

// PostBody ... Postリクエスト(URL, Body)
func (h *HTTP) PostBody(ctx context.Context, u string, body []byte, opt *HTTPOption) (bool, int, []byte) {
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(body))
	if err != nil {
		log.Warningf(ctx, "create request error: %s", err.Error())
		return false, 0, nil
	}
	for key, value := range opt.Headers {
		req.Header.Set(key, value)
	}
	return h.send(ctx, req, opt.Timeout)
}

func (h *HTTP) send(ctx context.Context, req *http.Request, timeout time.Duration) (bool, int, []byte) {
	dump, err := httputil.DumpRequestOut(req, true)
	if err == nil {
		log.Debugf(ctx, "send http request: %s", dump)
	} else {
		log.Warningf(ctx, "dumb http request error: %s, error=%s", spew.Sdump(req), err.Error())
	}

	client := urlfetch.Client(ctx)
	if timeout > 0 {
		client.Timeout = timeout
	} else {
		client.Timeout = h.Timeout
	}

	res, err := client.Do(req)
	if err != nil {
		log.Warningf(ctx, "http request error: %s", err.Error())
		return false, 0, nil
	}

	dump, err = httputil.DumpResponse(res, true)
	if err == nil {
		log.Debugf(ctx, "http response: %s", dump)
	} else {
		log.Warningf(ctx, "dumb http response error: %s, error=%s", spew.Sdump(req), err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warningf(ctx, "read http response body error: %s, error=%s", spew.Sdump(res), err.Error())
		return true, res.StatusCode, nil
	}
	defer res.Body.Close()

	return true, res.StatusCode, body
}

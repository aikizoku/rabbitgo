package httpclient

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

const defaultTimeout time.Duration = 15

// HTTPOption ... HTTP通信モジュールの追加設定
type HTTPOption struct {
	Headers map[string]string
	Timeout time.Duration
}

// Get ... Getリクエスト(URL)
func Get(ctx context.Context, u string, opt *HTTPOption) (int, []byte, error) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		log.Warningf(ctx, "create request error: %s", err.Error())
		return 0, nil, err
	}

	for key, value := range opt.Headers {
		req.Header.Set(key, value)
	}

	return send(ctx, req, opt.Timeout)
}

// GetForm ... Getリクエスト(URL, Params)
func GetForm(ctx context.Context, u string, params map[string]string, opt *HTTPOption) (int, []byte, error) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		log.Warningf(ctx, "create request error: %s", err.Error())
		return 0, nil, err
	}

	for key, value := range opt.Headers {
		req.Header.Set(key, value)
	}

	query := req.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}

	req.URL.RawQuery = query.Encode()
	return send(ctx, req, opt.Timeout)
}

// GetQueryString ... Getリクエスト(URL, QueryString)
func GetQueryString(ctx context.Context, u string, qs string, opt *HTTPOption) (int, []byte, error) {
	req, err := http.NewRequest("GET", u+"?"+qs, nil)
	if err != nil {
		log.Warningf(ctx, "create request error: %s", err.Error())
		return 0, nil, err
	}
	for key, value := range opt.Headers {
		req.Header.Set(key, value)
	}
	return send(ctx, req, opt.Timeout)
}

// PostForm ... Postリクエスト(URL, Params)
func PostForm(ctx context.Context, u string, params map[string]string, opt *HTTPOption) (int, []byte, error) {
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}
	req, err := http.NewRequest("POST", u, strings.NewReader(values.Encode()))
	if err != nil {
		log.Warningf(ctx, "create request error: %s", err.Error())
		return 0, nil, err
	}
	for key, value := range opt.Headers {
		req.Header.Set(key, value)
	}
	return send(ctx, req, opt.Timeout)
}

// PostJSON ... Postリクエスト(URL, JSON)
func PostJSON(ctx context.Context, u string, json []byte, opt *HTTPOption) (int, []byte, error) {
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(json))
	if err != nil {
		log.Warningf(ctx, "create request error: %s", err.Error())
		return 0, nil, err
	}

	for key, value := range opt.Headers {
		req.Header.Set(key, value)
	}
	opt.Headers["Content-Type"] = "application/json"

	return send(ctx, req, opt.Timeout)
}

// PostBody ... Postリクエスト(URL, Body)
func PostBody(ctx context.Context, u string, body []byte, opt *HTTPOption) (int, []byte, error) {
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(body))
	if err != nil {
		log.Warningf(ctx, "create request error: %s", err.Error())
		return 0, nil, err
	}

	for key, value := range opt.Headers {
		req.Header.Set(key, value)
	}

	return send(ctx, req, opt.Timeout)
}

func send(ctx context.Context, req *http.Request, timeout time.Duration) (int, []byte, error) {
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
		client.Timeout = defaultTimeout
	}

	res, err := client.Do(req)
	if err != nil {
		log.Warningf(ctx, "http request error: %s", err.Error())
		return 0, nil, err
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
		return res.StatusCode, nil, nil
	}
	defer res.Body.Close()

	return res.StatusCode, body, nil
}

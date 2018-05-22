package infrastructure

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

type HTTPRequestOption struct {
	Headers map[string]string
	Timeout time.Duration
}

func HTTPGet(ctx context.Context, url string, params map[string]string, opt *HTTPRequestOption) (bool, int, []byte) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
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

	dump, err := httputil.DumpRequestOut(req, true)
	if err == nil {
		log.Debugf(ctx, "send http request: %s", dump)
	} else {
		log.Errorf(ctx, "dumb http request: url=%s, params=%v, opt=%v", url, params, opt)
	}

	client := urlfetch.Client(ctx)
	if opt.Timeout > 0 {
		client.Timeout = opt.Timeout
	} else {
		client.Timeout = 10 * time.Second
	}

	res, err := client.Do(req)
	if err != nil {
		log.Errorf(ctx, "http request: error=%v", err)
		return false, 0, nil
	}

	dump, err = httputil.DumpResponse(res, true)
	if err == nil {
		log.Debugf(ctx, "http response: %s", dump)
	} else {
		log.Errorf(ctx, "dumb http response: url=%s, params=%v, opt=%v, response=%v", url, params, opt, res)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warningf(ctx, "unread body http response: url=%s, params=%v, opt=%v, response=%v", url, params, opt, res)
		return true, res.StatusCode, nil
	}
	defer res.Body.Close()

	return true, http.StatusOK, body
}

func HTTPPost() {

}

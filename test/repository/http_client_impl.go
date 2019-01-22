package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/aikizoku/skgo/src/lib/util"
)

const timeout time.Duration = 15 * time.Second

type httpClient struct {
}

func (r *httpClient) Get(u string, headers map[string]string) (int, string) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		panic(err)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return r.send(req)
}

func (r *httpClient) GetForm(u string, params map[string]interface{}, headers map[string]string) (int, string) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		panic(err.Error())
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	query := req.URL.Query()
	for key, value := range params {
		query.Add(key, value.(string))
	}
	req.URL.RawQuery = query.Encode()
	return r.send(req)
}

func (r *httpClient) PostJSON(url string, json []byte, headers map[string]string) (int, string) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		panic(err.Error())
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return r.send(req)
}

func (r *httpClient) PutJSON(url string, json []byte, headers map[string]string) (int, string) {
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(json))
	if err != nil {
		panic(err.Error())
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return r.send(req)
}

func (r *httpClient) DeleteJSON(url string, json []byte, headers map[string]string) (int, string) {
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(json))
	if err != nil {
		panic(err.Error())
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return r.send(req)
}

func (r *httpClient) send(req *http.Request) (int, string) {
	d, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		panic(err.Error())
	}
	reqDump := util.BytesToStr(d)
	fmt.Println(reqDump)

	client := http.Client{}
	client.Timeout = timeout

	res, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}

	d, err = httputil.DumpResponse(res, true)
	if err != nil {
		panic(err.Error())
	}
	resDump := util.BytesToStr(d)
	fmt.Println(resDump)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	out := new(bytes.Buffer)
	err = json.Indent(out, body, "", "    ")
	if err != nil {
		panic(err.Error())
	}

	return res.StatusCode, out.String()
}

// NewHTTPClient ... HTTPClientを作成する
func NewHTTPClient() HTTPClient {
	return &httpClient{}
}

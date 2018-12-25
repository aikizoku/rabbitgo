package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aikizoku/beego/test/config"
	"github.com/aikizoku/beego/test/model"
	"github.com/aikizoku/beego/test/repository"
)

const (
	apiTypeRest string = "REST"
)

type rest struct {
	hRepo repository.HTTPClient

	apis    []*model.API
	url     string
	headers map[string]string

	ovStagingURL    string
	ovProductionURL string
}

func (s *rest) Get(name string, uri string, params map[string]interface{}) {
	ov := &model.APIOverview{
		Type: apiTypeRest,
		URL: &model.APIOverviewURL{
			Staging:    s.ovStagingURL,
			Production: s.ovProductionURL,
		},
		URI: uri,
	}

	uri, ps := s.margeParams(uri, params)

	hs := s.createHeadersString(s.headers)

	qs := s.createQueryString(params)

	req := &model.APIRequest{
		Method:  http.MethodGet,
		URI:     uri,
		Headers: hs,
		Params:  qs,
	}

	status, body := s.hRepo.GetForm(s.url+uri, ps, s.headers)

	res := &model.APIResponse{
		StatusCode: status,
		Body:       body,
	}

	s.apis = append(s.apis, &model.API{
		Name:     name,
		Overview: ov,
		Request:  req,
		Response: res,
	})
}

func (s *rest) Post(name string, uri string, params map[string]interface{}) {
	ov := &model.APIOverview{
		Type: apiTypeRest,
		URL: &model.APIOverviewURL{
			Staging:    s.ovStagingURL,
			Production: s.ovProductionURL,
		},
		URI: uri,
	}

	uri, ps := s.margeParams(uri, params)

	s.headers["Content-Type"] = "application/json"
	hs := s.createHeadersString(s.headers)

	jPs, sPs := s.createJSON(ps)

	req := &model.APIRequest{
		Method:  http.MethodPost,
		URI:     uri,
		Headers: hs,
		Params:  sPs,
	}

	status, body := s.hRepo.PostJSON(s.url+uri, jPs, s.headers)

	res := &model.APIResponse{
		StatusCode: status,
		Body:       body,
	}

	s.apis = append(s.apis, &model.API{
		Name:     name,
		Overview: ov,
		Request:  req,
		Response: res,
	})
}

func (s *rest) Put(name string, uri string, params map[string]interface{}) {
	ov := &model.APIOverview{
		Type: apiTypeRest,
		URL: &model.APIOverviewURL{
			Staging:    s.ovStagingURL,
			Production: s.ovProductionURL,
		},
		URI: uri,
	}

	uri, ps := s.margeParams(uri, params)

	s.headers["Content-Type"] = "application/json"
	hs := s.createHeadersString(s.headers)

	jPs, sPs := s.createJSON(ps)

	req := &model.APIRequest{
		Method:  http.MethodPut,
		URI:     uri,
		Headers: hs,
		Params:  sPs,
	}

	status, body := s.hRepo.PutJSON(s.url+uri, jPs, s.headers)

	res := &model.APIResponse{
		StatusCode: status,
		Body:       body,
	}

	s.apis = append(s.apis, &model.API{
		Name:     name,
		Overview: ov,
		Request:  req,
		Response: res,
	})
}

func (s *rest) Delete(name string, uri string, params map[string]interface{}) {
	ov := &model.APIOverview{
		Type: apiTypeRest,
		URL: &model.APIOverviewURL{
			Staging:    s.ovStagingURL,
			Production: s.ovProductionURL,
		},
		URI: uri,
	}

	uri, ps := s.margeParams(uri, params)

	s.headers["Content-Type"] = "application/json"
	hs := s.createHeadersString(s.headers)

	jPs, sPs := s.createJSON(ps)

	req := &model.APIRequest{
		Method:  http.MethodDelete,
		URI:     uri,
		Headers: hs,
		Params:  sPs,
	}

	status, body := s.hRepo.DeleteJSON(s.url+uri, jPs, s.headers)

	res := &model.APIResponse{
		StatusCode: status,
		Body:       body,
	}

	s.apis = append(s.apis, &model.API{
		Name:     name,
		Overview: ov,
		Request:  req,
		Response: res,
	})
}

func (s *rest) GetAPIs() []*model.API {
	return s.apis
}

func (s *rest) margeParams(uri string, params map[string]interface{}) (string, map[string]interface{}) {
	ps := map[string]interface{}{}
	for key, value := range params {
		k := fmt.Sprintf("{%s}", key)
		if strings.Contains(uri, k) {
			uri = strings.Replace(uri, k, value.(string), -1)
		} else {
			ps[key] = value
		}
	}
	return uri, ps
}

func (s *rest) createQueryString(params map[string]interface{}) string {
	if len(params) == 0 {
		return ""
	}
	qs := "?"
	qsList := []string{}
	for key, value := range params {
		qsList = append(qsList, fmt.Sprintf("%s=%s", key, value))
	}
	qs += strings.Join(qsList, "&")
	return qs
}

func (s *rest) createHeadersString(headers map[string]string) string {
	hs := ""
	for key, value := range headers {
		if key == "Authorization" {
			// AuthorizationHeaderの値を隠蔽
			hs += fmt.Sprintf("%s: %s%s\n", key, config.AuthorizationPrefix, "XXXXXXXXXX")
		} else {
			hs += fmt.Sprintf("%s: %s%s\n", key, config.AuthorizationPrefix, value)
		}
	}
	return hs
}

func (s *rest) createJSON(params map[string]interface{}) ([]byte, string) {
	dstJSON, err := json.Marshal(params)
	if err != nil {
		panic(err.Error())
	}

	out := new(bytes.Buffer)
	err = json.Indent(out, dstJSON, "", "    ")
	if err != nil {
		panic(err.Error())
	}
	dstStr := out.String()

	return dstJSON, dstStr
}

// NewRest ... Restを作成する
func NewRest(
	hRepo repository.HTTPClient,
	url string,
	headers map[string]string,
	stagingURL string,
	productionURL string) Rest {
	return &rest{
		hRepo:           hRepo,
		apis:            []*model.API{},
		url:             url,
		headers:         headers,
		ovStagingURL:    stagingURL,
		ovProductionURL: productionURL,
	}
}

package repository

import (
	"bytes"

	"github.com/alecthomas/template"
)

type templateClient struct {
}

// GetMarged ... 任意の値をマージした文字列を返す
func (r *templateClient) GetMarged(path string, src interface{}) string {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}

	var doc bytes.Buffer
	if err := tmpl.Execute(&doc, src); err != nil {
		panic(err)
	}
	return doc.String()
}

// NewTemplateClient ... TemplateClientを作成する
func NewTemplateClient() TemplateClient {
	return &templateClient{}
}

package environment

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/deploy"
)

// Load ... 環境変数を読み込む
func Load() {
	// 値
	file, err := ioutil.ReadFile("./env.yaml")
	if err != nil {
		panic(err)
	}
	val := &Variable{}
	err = yaml.Unmarshal(file, &val)
	if err != nil {
		panic(err)
	}

	var src map[string]string
	if deploy.IsLocal() {
		file, err := ioutil.ReadFile("../../project.json")
		if err != nil {
			panic(err)
		}
		prj := &Project{}
		err = json.Unmarshal(file, &prj)
		if err != nil {
			panic(err)
		}
		src = val.Local
		src["PROJECT_ID"] = prj.Local
		src["DEPLOY"] = "local"
	} else if deploy.IsStaging() {
		prj := os.Getenv("GOOGLE_CLOUD_PROJECT")
		if prj == "" {
			panic("env not found GOOGLE_CLOUD_PROJECT")
		}
		src = val.Staging
		src["PROJECT_ID"] = prj
	} else if deploy.IsProduction() {
		prj := os.Getenv("GOOGLE_CLOUD_PROJECT")
		if prj == "" {
			panic("env not found GOOGLE_CLOUD_PROJECT")
		}
		src = val.Production
		src["PROJECT_ID"] = prj
	} else {
		panic(fmt.Errorf("invalid deploy: %s", os.Getenv("DEPLOY")))
	}

	for k, v := range src {
		err = os.Setenv(k, v)
		if err != nil {
			panic(err)
		}
	}
}

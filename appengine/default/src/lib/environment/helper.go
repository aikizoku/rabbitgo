package environment

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/deploy"
)

// Load ... 環境変数を読み込む
func Load() {
	// プロジェクト
	file, err := ioutil.ReadFile("../../project.json")
	if err != nil {
		panic(err)
	}
	prj := &Project{}
	err = json.Unmarshal(file, &prj)
	if err != nil {
		panic(err)
	}

	// 値
	file, err = ioutil.ReadFile("./env.json")
	if err != nil {
		panic(err)
	}
	val := &Variable{}
	err = json.Unmarshal(file, &val)
	if err != nil {
		panic(err)
	}

	var src map[string]string
	if deploy.IsLocal() {
		src = val.Local
		src["PROJECT_ID"] = prj.Local
		src["DEPLOY"] = "local"
	} else if deploy.IsStaging() {
		src = val.Staging
		src["PROJECT_ID"] = prj.Staging
		src["DEPLOY"] = "staging"
	} else if deploy.IsProduction() {
		src = val.Production
		src["PROJECT_ID"] = prj.Production
		src["DEPLOY"] = "production"
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

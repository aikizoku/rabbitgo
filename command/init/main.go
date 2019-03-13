package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rabee-inc/push/command/common"
	"gopkg.in/yaml.v2"
)

func main() {
	// env.jsonの読み込み
	env := common.LoadEnvFile()

	// ProjectIDsの読み込み
	pIDs := common.GetProjectIDs(env)

	// リセット
	os.RemoveAll("./deploy")

	// 初期化
	for _, app := range env.Apps {
		createDeployDir(common.Local, app)
		createHotReloadLinks(common.Local, app)
		createValuesFile(common.Local, app, pIDs.Local, env.Values.Local)
		createCredentialsFile(common.Local, app, env.Credentials.Local)

		createDeployDir(common.Staging, app)
		createHotReloadLinks(common.Staging, app)
		createValuesFile(common.Staging, app, pIDs.Staging, env.Values.Staging)
		createCredentialsFile(common.Staging, app, env.Credentials.Staging)

		createDeployDir(common.Production, app)
		createHotReloadLinks(common.Production, app)
		createValuesFile(common.Production, app, pIDs.Production, env.Values.Production)
		createCredentialsFile(common.Production, app, env.Credentials.Production)
	}
}

func createDeployDir(env string, app string) {
	os.MkdirAll(fmt.Sprintf("./deploy/appengine/%s/%s", env, app), 0755)
}

func createHotReloadLinks(env string, app string) {
	// app.yaml
	os.Symlink(
		fmt.Sprintf("../../../../appengine/app/%s/app_%s.yaml", app, env),
		fmt.Sprintf("deploy/appengine/%s/%s/app.yaml", env, app))
	// go
	os.Symlink(
		fmt.Sprintf("../../../../appengine/app/%s/main.go", app),
		fmt.Sprintf("deploy/appengine/%s/%s/main.go", env, app))
	os.Symlink(
		fmt.Sprintf("../../../../appengine/app/%s/dependency.go", app),
		fmt.Sprintf("deploy/appengine/%s/%s/dependency.go", env, app))
	os.Symlink(
		fmt.Sprintf("../../../../appengine/app/%s/routing.go", app),
		fmt.Sprintf("deploy/appengine/%s/%s/routing.go", env, app))
	// cron.yaml
	os.Symlink(
		fmt.Sprintf("../../../../appengine/config/cron.yaml"),
		fmt.Sprintf("deploy/appengine/%s/%s/cron.yaml", env, app))
	os.Symlink(
		fmt.Sprintf("../../../appengine/config/cron.yaml"),
		fmt.Sprintf("deploy/appengine/%s/cron.yaml", env))
	// dispatch.yaml
	os.Symlink(
		fmt.Sprintf("../../../../appengine/config/dispatch_%s.yaml", env),
		fmt.Sprintf("deploy/appengine/%s/%s/dispatch.yaml", env, app))
	os.Symlink(
		fmt.Sprintf("../../../appengine/config/dispatch_%s.yaml", env),
		fmt.Sprintf("deploy/appengine/%s/dispatch.yaml", env))
	// index.yaml
	os.Symlink(
		fmt.Sprintf("../../../../appengine/config/index.yaml"),
		fmt.Sprintf("deploy/appengine/%s/%s/index.yaml", env, app))
	os.Symlink(
		fmt.Sprintf("../../../appengine/config/index.yaml"),
		fmt.Sprintf("deploy/appengine/%s/index.yaml", env))
	// queue.yaml
	os.Symlink(
		fmt.Sprintf("../../../../appengine/config/queue.yaml"),
		fmt.Sprintf("deploy/appengine/%s/%s/queue.yaml", env, app))
	os.Symlink(
		fmt.Sprintf("../../../appengine/config/queue.yaml"),
		fmt.Sprintf("deploy/appengine/%s/queue.yaml", env))
	// src
	os.Symlink(
		fmt.Sprintf("../../../../src"),
		fmt.Sprintf("deploy/appengine/%s/%s/src", env, app))
	// .gcloudignore
	os.Symlink(
		fmt.Sprintf("../../../../.gcloudignore"),
		fmt.Sprintf("deploy/appengine/%s/%s/.gcloudignore", env, app))
}

func createValuesFile(env string, app string, pID string, data map[string]string) {
	yData := map[string]interface{}{}
	data["PROJECT_ID"] = pID
	data["ENV"] = env
	data["GOOGLE_APPLICATION_CREDENTIALS"] = "./credentials.json"
	yData["env_variables"] = data
	y, err := yaml.Marshal(yData)
	if err != nil {
		panic(err.Error())
	}
	common.CreateFile(
		fmt.Sprintf("./deploy/appengine/%s/%s/values.yaml", env, app),
		string(y),
	)
}

func createCredentialsFile(env string, app string, data map[string]string) {
	j, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}
	common.CreateFile(
		fmt.Sprintf("./deploy/appengine/%s/%s/credentials.json", env, app),
		string(j),
	)
}

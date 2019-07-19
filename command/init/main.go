package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/aikizoku/rabbitgo/command/common"
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

func createDeployDir(deploy string, app string) {
	os.MkdirAll(fmt.Sprintf("./deploy/%s/%s", deploy, app), 0755)
}

func createHotReloadLinks(deploy string, app string) {
	// app.yaml
	if deploy != common.Local {
		os.Symlink(
			fmt.Sprintf("../../../%s/app_%s.yaml", app, deploy),
			fmt.Sprintf("deploy/%s/%s/app.yaml", deploy, app))
	}

	// go
	os.Symlink(
		fmt.Sprintf("../../../%s/enviroment.go", app),
		fmt.Sprintf("deploy/%s/%s/enviroment.go", deploy, app))
	os.Symlink(
		fmt.Sprintf("../../../%s/dependency.go", app),
		fmt.Sprintf("deploy/%s/%s/dependency.go", deploy, app))
	os.Symlink(
		fmt.Sprintf("../../../%s/routing.go", app),
		fmt.Sprintf("deploy/%s/%s/routing.go", deploy, app))
	os.Symlink(
		fmt.Sprintf("../../../%s/main.go", app),
		fmt.Sprintf("deploy/%s/%s/main.go", deploy, app))

	// go mod
	goModFile := fmt.Sprintf("./deploy/%s/%s/go.mod", deploy, app)
	common.ExecCommand(
		"cp",
		fmt.Sprintf("./%s/go.mod", app),
		goModFile,
	)
	common.ReplaceFile(goModFile, "../src", "./src")
	os.Symlink(
		fmt.Sprintf("../../../%s/go.sum", app),
		fmt.Sprintf("deploy/%s/%s/go.sum", deploy, app))

	// src
	os.Symlink(
		fmt.Sprintf("../../../src"),
		fmt.Sprintf("deploy/%s/%s/src", deploy, app))

	// .gcloudignore
	os.Symlink(
		fmt.Sprintf("../../../../.gcloudignore"),
		fmt.Sprintf("deploy/%s/%s/.gcloudignore", deploy, app))
}

func createValuesFile(deploy string, app string, pID string, data map[string]string) {
	data["PROJECT_ID"] = pID
	data["DEPLOY"] = deploy
	data["GOOGLE_APPLICATION_CREDENTIALS"] = "./credentials.json"
	y, err := yaml.Marshal(data)
	if err != nil {
		panic(err.Error())
	}
	common.CreateFile(fmt.Sprintf(
		"./deploy/%s/%s/.env", deploy, app),
		string(y),
	)
}

func createCredentialsFile(deploy string, app string, data map[string]string) {
	j, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}
	common.CreateFile(
		fmt.Sprintf("./deploy/%s/%s/credentials.json", deploy, app),
		string(j),
	)
}

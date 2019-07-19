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
	createDeployDir(common.Local)
	createSourceFile(common.Local)
	createEnvironmentFile(common.Local, pIDs.Local, env.Appengine.Local)
	createCredentialsFile(common.Local, env.Credentials.Local)

	createDeployDir(common.Staging)
	createSourceFile(common.Staging)
	createEnvironmentFile(common.Staging, pIDs.Staging, env.Appengine.Staging)
	createCredentialsFile(common.Staging, env.Credentials.Staging)

	createDeployDir(common.Production)
	createSourceFile(common.Production)
	createEnvironmentFile(common.Production, pIDs.Production, env.Appengine.Production)
	createCredentialsFile(common.Production, env.Credentials.Production)
}

func createDeployDir(deploy string) {
	os.MkdirAll(fmt.Sprintf("./deploy/%s", deploy), 0755)
}

func createSourceFile(deploy string) {
	// app.yaml
	if deploy != common.Local {
		os.Symlink(
			fmt.Sprintf("../../app_%s.yaml", deploy),
			fmt.Sprintf("deploy/%s/app.yaml", deploy))
	}

	// go
	os.Symlink(
		"../../enviroment.go",
		fmt.Sprintf("deploy/%s/enviroment.go", deploy))
	os.Symlink(
		"../../dependency.go",
		fmt.Sprintf("deploy/%s/dependency.go", deploy))
	os.Symlink(
		"../../routing.go",
		fmt.Sprintf("deploy/%s/routing.go", deploy))
	os.Symlink(
		"../../main.go",
		fmt.Sprintf("deploy/%s/main.go", deploy))

	// go mod
	os.Symlink(
		"../../go.mod",
		fmt.Sprintf("deploy/%s/go.mod", deploy))
	os.Symlink(
		"../../go.sum",
		fmt.Sprintf("deploy/%s/go.sum", deploy))

	// src
	os.Symlink(
		"../../src",
		fmt.Sprintf("deploy/%s/src", deploy))

	// .gcloudignore
	os.Symlink(
		"../../../.gcloudignore",
		fmt.Sprintf("deploy/%s/.gcloudignore", deploy))
}

func createEnvironmentFile(deploy string, pID string, data map[string]interface{}) {
	data["PROJECT_ID"] = pID
	data["DEPLOY"] = deploy
	data["GOOGLE_APPLICATION_CREDENTIALS"] = "./credentials.json"
	y, err := yaml.Marshal(data)
	if err != nil {
		panic(err.Error())
	}
	common.CreateFile(fmt.Sprintf(
		"./deploy/%s/.env", deploy),
		string(y),
	)
}

func createCredentialsFile(deploy string, data map[string]interface{}) {
	j, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}
	common.CreateFile(
		fmt.Sprintf("./deploy/%s/credentials.json", deploy),
		string(j),
	)
}

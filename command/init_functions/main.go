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
	createDeployDir(common.Staging)
	for app, config := range env.Functions.Staging {
		createSourceFile(common.Staging, app)
		createEnvironmentFile(common.Staging, app, pIDs.Staging, config.(map[string]interface{})["environment"].(map[string]interface{}))
		createCredentialsFile(common.Staging, app, env.Credentials.Staging)
	}

	createDeployDir(common.Production)
	for app, config := range env.Functions.Production {
		createSourceFile(common.Production, app)
		createEnvironmentFile(common.Production, app, pIDs.Production, config.(map[string]interface{})["environment"].(map[string]interface{}))
		createCredentialsFile(common.Production, app, env.Credentials.Production)
	}
}

func createDeployDir(deploy string) {
	os.MkdirAll(fmt.Sprintf("./deploy/%s", deploy), 0755)
}

func createSourceFile(deploy string, app string) {
	// src
	common.ExecCommand(
		"cp",
		"-r",
		fmt.Sprintf("./%s", app),
		fmt.Sprintf("./deploy/%s/%s", deploy, app))

	// .gcloudignore
	os.Symlink(
		"../../../../.gcloudignore",
		fmt.Sprintf("deploy/%s/%s/.gcloudignore", deploy, app))
}

func createEnvironmentFile(deploy string, app string, pID string, data map[string]interface{}) {
	data["PROJECT_ID"] = pID
	data["DEPLOY"] = deploy
	data["GOOGLE_APPLICATION_CREDENTIALS"] = "./credentials.json"
	y, err := yaml.Marshal(data)
	if err != nil {
		panic(err.Error())
	}
	common.CreateFile(fmt.Sprintf(
		"./deploy/%s/%s/env.yaml", deploy, app),
		string(y),
	)
}

func createCredentialsFile(deploy string, app string, data map[string]interface{}) {
	j, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}
	common.CreateFile(
		fmt.Sprintf("./deploy/%s/%s/credentials.json", deploy, app),
		string(j),
	)
}

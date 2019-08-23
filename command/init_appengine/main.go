package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/aikizoku/rabbitgo/command/common"
)

func main() {
	// リセット
	os.RemoveAll("./deploy")

	deploys := []string{common.Local, common.Staging, common.Production}

	for _, deploy := range deploys {
		env := common.LoadEnvFile(deploy)
		for _, app := range env.Apps {
			createDeployDir(deploy, app)
			createSourceFile(deploy, app)
			createEnvironmentFile(deploy, app, env.GetProjectID(), env.Appengine)
			createCredentialsFile(deploy, app, env.Credentials)
		}
	}
}

func createDeployDir(deploy string, app string) {
	os.MkdirAll(fmt.Sprintf("./deploy/%s/%s", deploy, app), 0755)
}

func createSourceFile(deploy string, app string) {
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

	// realize
	common.ExecCommand(
		"cp",
		fmt.Sprintf("./%s/.realize.yaml", app),
		fmt.Sprintf("./deploy/%s/%s/.realize.yaml", deploy, app),
	)

	// .gcloudignore
	os.Symlink(
		fmt.Sprintf("../../../../.gcloudignore"),
		fmt.Sprintf("deploy/%s/%s/.gcloudignore", deploy, app))
}

func createEnvironmentFile(deploy string, app string, pID string, data map[string]interface{}) {
	data["PROJECT_ID"] = pID
	data["SERVICE_ID"] = app
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

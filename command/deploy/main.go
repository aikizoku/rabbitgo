package main

import (
	"flag"
	"fmt"

	"github.com/aikizoku/merlin/command/common"
)

func main() {
	var (
		env = flag.String("env", "local", "environment")
		app = flag.String("app", "api", "appengine application name")
	)
	flag.Parse()

	// env.jsonの読み込み
	e := common.LoadEnvFile()

	// ProjectIDを取得
	pIDs := common.GetProjectIDs(e)

	common.ExecCommand(
		"gcloud",
		"app",
		"deploy",
		"-q",
		fmt.Sprintf("deploy/appengine/%s/%s/app.yaml", *env, *app),
		"--project",
		pIDs.GetByEnv(*env),
	)
}

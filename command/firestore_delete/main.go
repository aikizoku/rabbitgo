package main

import (
	"flag"

	"github.com/aikizoku/merlin/command/common"
)

func main() {
	var (
		env = flag.String("env", "local", "environment")
	)
	flag.Parse()

	// env.jsonの読み込み
	e := common.LoadEnvFile()

	// ProjectIDを取得
	pIDs := common.GetProjectIDs(e)

	common.ExecCommand(
		"firebase",
		"firestore:delete",
		"-y",
		"--all-collections",
		"--project",
		pIDs.GetByEnv(*env),
	)
}

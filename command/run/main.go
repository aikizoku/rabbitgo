package main

import (
	"flag"
	"fmt"

	"github.com/rabee-inc/push/command/common"
)

func main() {
	var (
		env = flag.String("env", "local", "environment")
		app = flag.String("app", "api", "appengine application name")
	)
	flag.Parse()

	common.ExecCommand(
		"dev_appserver.py",
		fmt.Sprintf("deploy/appengine/%s/%s/app.yaml", *env, *app),
	)
}

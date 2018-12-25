package main

import (
	"flag"

	"github.com/aikizoku/beego/test/config"
)

func main() {
	// Args
	snro := flag.String("scenario", "normal", "test scenario name")
	url := flag.String("url", "http://localhost:8080", "api endpoint url")
	auth := flag.String("auth", "", "authorization header value")
	flag.Parse()

	// Dependency
	d := &config.Dependency{}
	d.Inject(*snro, *url, *auth)

	// Execute
	d.Scenario.Run()
}

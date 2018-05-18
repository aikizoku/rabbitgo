package main

import (
	"flag"
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var (
	env        = flag.String("e", "development", "app server enviroment")
	configFile = flag.String("c", "config.yaml", "app server configuration file")
	basePath   = flag.String("b", ".", "app server current path")
)

func main() {
	flag.Parse()
	fmt.Printf("env: %s", *env)
	fmt.Printf("configFile: %s", *configFile)
	fmt.Printf("basePath: %s", *basePath)
	http.HandleFunc("/", handler)
	appengine.Main()
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "debug_log: %s", "Hello")
	fmt.Fprintln(w, "Hello, world!")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("pong"))
}

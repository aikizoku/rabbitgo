package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/aikizoku/go-gae-template/src/handler/api"
	"github.com/aikizoku/go-gae-template/src/middleware"
	"github.com/aikizoku/go-gae-template/src/repository"
	"github.com/aikizoku/go-gae-template/src/service"
	"github.com/go-chi/chi"
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

	r := chi.NewRouter()

	// Setup Middleware

	// Dependency Injection
	sampleRepo := repository.NewSample()
	sampleSvc := service.NewSample(sampleRepo)
	sampleHandler := &api.SampleHandler{
		Service: sampleSvc,
	}

	// Routing
	rpc := *middleware.NewJsonrpc2()
	rpc.Register("sample", sampleHandler)

	jsonrpc2 := api.Jsonrpc2{
		Rpc: rpc,
	}
	r.Post("/api/v1/rpc", jsonrpc2.Handler)

	// Run
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

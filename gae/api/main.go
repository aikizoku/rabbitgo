package main

import (
	"net/http"
	"time"

	"github.com/aikizoku/go-gae-template/src/config"
	"github.com/aikizoku/go-gae-template/src/handler"
	"github.com/aikizoku/go-gae-template/src/handler/api"
	"github.com/aikizoku/go-gae-template/src/infrastructure"
	"github.com/aikizoku/go-gae-template/src/middleware"
	"github.com/aikizoku/go-gae-template/src/repository"
	"github.com/aikizoku/go-gae-template/src/service"
	"github.com/go-chi/chi"
	"google.golang.org/appengine"
)

func main() {
	r := chi.NewRouter()

	// Dependency Injection
	csql := infrastructure.NewCSQLClient(config.GetCSQLConfig("trial"))

	sampleRepo := repository.NewSample(infrastructure.NewHTTP(10*time.Second), csql)

	sampleSvc := service.NewSample(sampleRepo)

	sampleHandler := &api.SampleHandler{
		Svc: sampleSvc,
	}

	// Routing
	r.Use(middleware.AccessControl)
	r.Get("/ping", handler.PingHandler)

	rpc := *middleware.NewJsonrpc2()
	r.Route("/v1/rpc", func(subr chi.Router) {
		subr.Use(rpc.Handle)
		subr.Post("/", func(w http.ResponseWriter, r *http.Request) {})
	})

	http.Handle("/", r)

	// API
	rpc.Register("sample", sampleHandler)

	// Run
	appengine.Main()
}

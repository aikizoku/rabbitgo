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
	defer func() {
		csql.Close()
	}()

	sampleRepo := repository.NewSample(infrastructure.NewHTTP(10 * time.Second))

	sampleSvc := service.NewSample(sampleRepo)

	sampleHandler := &api.SampleHandler{
		Svc: sampleSvc,
	}
	testPutHandler := &api.TestPutHandler{
		Svc: sampleSvc,
	}
	testGetHandler := &api.TestGetHandler{
		Svc: sampleSvc,
	}
	testDeleteHandler := &api.TestDeleteHandler{
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
	rpc.Register("test_put", testPutHandler)
	rpc.Register("test_get", testGetHandler)
	rpc.Register("test_delete", testDeleteHandler)

	// Run
	appengine.Main()
}

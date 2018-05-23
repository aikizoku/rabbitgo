package main

import (
	"net/http"
	"time"

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
	sampleRepo := repository.NewSample(infrastructure.NewHTTP(10 * time.Second))
	sampleSvc := service.NewSample(sampleRepo)
	sampleHandler := &api.SampleHandler{
		Svc: sampleSvc,
	}

	// Routing
	// r.Use(middleware.BasicAuth)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})
	rpc := *middleware.NewJsonrpc2()
	r.Route("/v1/rpc", func(subr chi.Router) {
		subr.Use(rpc.Handle)
	})
	http.Handle("/", r)

	// API
	rpc.Register("sample", sampleHandler)

	// Run
	appengine.Main()
}

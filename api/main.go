package main

import (
	"net/http"

	"github.com/aikizoku/go-gae-template/src/handler/api"
	"github.com/aikizoku/go-gae-template/src/middleware"
	"github.com/aikizoku/go-gae-template/src/repository"
	"github.com/aikizoku/go-gae-template/src/service"
	"github.com/go-chi/chi"
	"google.golang.org/appengine"
)

func main() {
	r := chi.NewRouter()

	// Setup Middleware

	// Dependency Injection
	sampleRepo := repository.NewSample()
	sampleSvc := service.NewSample(sampleRepo)
	sampleHandler := &api.SampleHandler{
		Svc: sampleSvc,
	}

	// Routing
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	rpc := *middleware.NewJsonrpc2()
	rpc.Register("sample", sampleHandler)

	jsonrpc2 := api.Jsonrpc2{
		Rpc: rpc,
	}
	r.Post("/api/v1/rpc", jsonrpc2.Handler)

	http.Handle("/", r)

	// Run
	appengine.Main()
}

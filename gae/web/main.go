package main

import (
	"github.com/aikizoku/go-gae-template/src/handler/web"
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
	sampleWeb := &web.Sample{
		service: sampleSvc,
	}

	// Routing
	sampleWeb.Routing(r)

	// Run
	appengine.Main()
}

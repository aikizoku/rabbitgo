package main

import (
	"net/http"
	"time"

	"github.com/aikizoku/go-gae-template/src/handler"
	"github.com/aikizoku/go-gae-template/src/handler/task"
	"github.com/aikizoku/go-gae-template/src/infrastructure"

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

	sampleRgs := task.SampleRegister{}
	sampleWkr := task.SampleWorker{Svc: sampleSvc}

	// Routing
	r.Get("/ping", handler.PingHandler)

	r.Route("/register", func(subr chi.Router) {
		subr.Get("/sample", sampleRgs.HogeRegister)
	})

	r.Route("/worker", func(subr chi.Router) {
		subr.Post("/sample", sampleWkr.HogeWorker)
	})

	http.Handle("/", r)

	// Run
	appengine.Main()
}

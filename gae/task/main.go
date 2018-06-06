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

	sampleQue := task.SampleQueueing{}
	sampleTsk := task.SampleTask{Svc: sampleSvc}

	// Routing
	r.Get("/ping", handler.PingHandler)

	r.Route("/queueing", func(subr chi.Router) {
		subr.Get("/sample", sampleQue.HogeQueueing)
	})

	r.Route("/task", func(subr chi.Router) {
		subr.Post("/sample", sampleTsk.HogeTask)
	})

	http.Handle("/", r)

	// Run
	appengine.Main()
}

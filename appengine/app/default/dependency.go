package main

import (
	"os"

	"github.com/aikizoku/merlin/src/handler/api"
	"github.com/aikizoku/merlin/src/lib/firebaseauth"
	"github.com/aikizoku/merlin/src/lib/httpheader"
	"github.com/aikizoku/merlin/src/lib/jsonrpc2"
	"github.com/aikizoku/merlin/src/repository"
	"github.com/aikizoku/merlin/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	FirebaseAuth    *firebaseauth.Middleware
	HTTPHeader      *httpheader.Middleware
	JSONRPC2Handler *jsonrpc2.Handler
	SampleHandler   *api.SampleHandler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject() {
	// Config
	crePath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if crePath == "" {
		panic("no config error: GOOGLE_APPLICATION_CREDENTIALS")
	}

	// Repository
	repo := repository.NewSample()

	// Service
	faSvc := firebaseauth.NewService()
	hhSvc := httpheader.NewService()
	svc := service.NewSample(repo)

	// Middleware
	d.FirebaseAuth = firebaseauth.NewMiddleware(faSvc)
	d.HTTPHeader = httpheader.NewMiddleware(hhSvc)

	// Handler
	d.JSONRPC2Handler = jsonrpc2.NewHandler()
	d.SampleHandler = api.NewSampleHandler(svc)
}

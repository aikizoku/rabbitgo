package main

import (
	"os"

	"github.com/aikizoku/skgo/src/handler/api"
	"github.com/aikizoku/skgo/src/lib/cloudfirestore"
	"github.com/aikizoku/skgo/src/lib/firebaseauth"
	"github.com/aikizoku/skgo/src/lib/httpheader"
	"github.com/aikizoku/skgo/src/lib/jsonrpc2"
	"github.com/aikizoku/skgo/src/repository"
	"github.com/aikizoku/skgo/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	DummyFirebaseAuth *firebaseauth.Middleware
	FirebaseAuth      *firebaseauth.Middleware
	DummyHTTPHeader   *httpheader.Middleware
	HTTPHeader        *httpheader.Middleware
	JSONRPC2          *jsonrpc2.Middleware
	SampleHandler     *api.SampleHandler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject() {
	// Config
	crePath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if crePath == "" {
		panic("no config error: GOOGLE_APPLICATION_CREDENTIALS")
	}

	// Client
	fCli, err := cloudfirestore.NewClient(crePath)
	if err != nil {
		panic(err.Error())
	}

	// Repository
	repo := repository.NewSample(fCli)

	// Service
	dfaSvc := firebaseauth.NewDummyService()
	faSvc := firebaseauth.NewService()
	dhhSvc := httpheader.NewDummyService()
	hhSvc := httpheader.NewService()
	svc := service.NewSample(repo)

	// Middleware
	d.DummyFirebaseAuth = firebaseauth.NewMiddleware(dfaSvc)
	d.FirebaseAuth = firebaseauth.NewMiddleware(faSvc)
	d.DummyHTTPHeader = httpheader.NewMiddleware(dhhSvc)
	d.HTTPHeader = httpheader.NewMiddleware(hhSvc)

	// Handler
	d.SampleHandler = api.NewSampleHandler(svc)
}

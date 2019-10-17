package main

import (
	"github.com/aikizoku/rabbitgo/appengine/src/handler/api"
	"github.com/aikizoku/rabbitgo/appengine/src/lib/cloudfirestore"
	"github.com/aikizoku/rabbitgo/appengine/src/lib/cloudpubsub"
	"github.com/aikizoku/rabbitgo/appengine/src/lib/deploy"
	"github.com/aikizoku/rabbitgo/appengine/src/lib/firebaseauth"
	"github.com/aikizoku/rabbitgo/appengine/src/lib/images"
	"github.com/aikizoku/rabbitgo/appengine/src/lib/jsonrpc2"
	"github.com/aikizoku/rabbitgo/appengine/src/lib/log"
	"github.com/aikizoku/rabbitgo/appengine/src/repository"
	"github.com/aikizoku/rabbitgo/appengine/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	Log             *log.Middleware
	FirebaseAuth    *firebaseauth.Middleware
	SampleHandler   *api.SampleHandler
	JSONRPC2Handler *jsonrpc2.Handler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject(e *Environment) {
	// Client
	aCli := firebaseauth.NewClient(e.CredentialsPath)
	fCli := cloudfirestore.NewClient(e.CredentialsPath)
	var lCli log.Writer
	if deploy.IsLocal() {
		lCli = log.NewWriterStdout()
	} else {
		lCli = log.NewWriterStackdriver(e.ProjectID)
	}
	psCli := cloudpubsub.NewClient(e.ProjectID, e.CredentialsPath, []string{"image-converter"})
	imgCli := images.NewClient(psCli, "image-converter")

	// Repository
	repo := repository.NewSample(fCli, imgCli)

	// Service
	var faSvc firebaseauth.Service
	if deploy.IsProduction() {
		faSvc = firebaseauth.NewService(aCli)
	} else {
		faSvc = firebaseauth.NewDebugService(aCli)
	}
	svc := service.NewSample(repo)

	// Middleware
	d.Log = log.NewMiddleware(lCli, e.MinLogSeverity)
	d.FirebaseAuth = firebaseauth.NewMiddleware(faSvc)

	// Handler
	d.SampleHandler = api.NewSampleHandler(svc)
	d.JSONRPC2Handler = jsonrpc2.NewHandler()
}

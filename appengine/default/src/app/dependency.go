package app

import (
	"github.com/rabee-inc/go-pkg/accesscontrol"
	"github.com/rabee-inc/go-pkg/cloudfirestore"
	"github.com/rabee-inc/go-pkg/cloudpubsub"
	"github.com/rabee-inc/go-pkg/deploy"
	"github.com/rabee-inc/go-pkg/firebaseauth"
	"github.com/rabee-inc/go-pkg/images"
	"github.com/rabee-inc/go-pkg/jsonrpc2"
	"github.com/rabee-inc/go-pkg/log"

	"github.com/aikizoku/rabbitgo/appengine/default/src/handler/api"
	"github.com/aikizoku/rabbitgo/appengine/default/src/repository"
	"github.com/aikizoku/rabbitgo/appengine/default/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	Accesscontrol   *accesscontrol.Middleware
	Log             *log.Middleware
	FirebaseAuth    *firebaseauth.Middleware
	SampleHandler   *api.SampleHandler
	JSONRPC2Handler *jsonrpc2.Handler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject(e *Environment) {
	// Client
	aCli := firebaseauth.NewClient(e.ProjectID)
	fCli := cloudfirestore.NewClient(e.ProjectID)
	var lCli log.Writer
	if deploy.IsLocal() {
		lCli = log.NewWriterStdout()
	} else {
		lCli = log.NewWriterStackdriver(e.ProjectID)
	}
	psCli := cloudpubsub.NewClient(e.ProjectID, []string{"image-converter"})
	imgCli := images.NewClient(psCli, "image-converter")

	// Repository
	repo := repository.NewSample(fCli, imgCli)

	// Service
	var faSvc firebaseauth.Service
	if deploy.IsProduction() {
		faSvc = firebaseauth.NewService(aCli)
	} else {
		faSvc = firebaseauth.NewServiceDebug(aCli, map[string]interface{}{})
	}
	svc := service.NewSample(repo)

	// Middleware
	d.Accesscontrol = accesscontrol.NewMiddleware(nil)
	d.Log = log.NewMiddleware(lCli, e.MinLogSeverity)
	d.FirebaseAuth = firebaseauth.NewMiddleware(faSvc)

	// Handler
	d.SampleHandler = api.NewSampleHandler(svc)
	d.JSONRPC2Handler = jsonrpc2.NewHandler()
}

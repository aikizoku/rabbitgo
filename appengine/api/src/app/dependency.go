package app

import (
	"github.com/rabee-inc/go-pkg/accesscontrol"
	"github.com/rabee-inc/go-pkg/cloudfirestore"
	"github.com/rabee-inc/go-pkg/cloudpubsub"
	"github.com/rabee-inc/go-pkg/deploy"
	"github.com/rabee-inc/go-pkg/firebaseauth"
	"github.com/rabee-inc/go-pkg/images"
	"github.com/rabee-inc/go-pkg/internalauth"
	"github.com/rabee-inc/go-pkg/log"

	"github.com/aikizoku/rabbitgo/appengine/api/src/handler/site"
	"github.com/aikizoku/rabbitgo/appengine/api/src/repository"
	"github.com/aikizoku/rabbitgo/appengine/api/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	Accesscontrol *accesscontrol.Middleware
	Log           *log.Middleware
	FirebaseAuth  *firebaseauth.Middleware
	InternalAuth  *internalauth.Middleware

	SampleHandler *site.SampleHandler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject(e *Environment) {
	// Client
	authCli := firebaseauth.NewClient(e.ProjectID)
	fCli := cloudfirestore.NewClient(e.ProjectID)
	var lCli log.Writer
	if deploy.IsLocal() {
		lCli = log.NewWriterStdout()
	} else {
		lCli = log.NewWriterStackdriver(e.ProjectID)
	}
	psCli := cloudpubsub.NewClient(e.ProjectID, []string{
		images.ConverterTopicID,
	})
	imgCli := images.NewClient(psCli)

	// Repository
	repo := repository.NewSample(fCli, imgCli)

	// Service
	var faSvc firebaseauth.Service
	if deploy.IsProduction() {
		faSvc = firebaseauth.NewService(authCli)
	} else {
		faSvc = firebaseauth.NewServiceDebug(authCli, map[string]interface{}{})
	}
	svc := service.NewSample(repo)

	// Middleware
	d.Accesscontrol = accesscontrol.NewMiddleware(nil)
	d.Log = log.NewMiddleware(lCli, e.MinLogSeverity)
	d.FirebaseAuth = firebaseauth.NewMiddleware(faSvc, false)
	d.InternalAuth = internalauth.NewMiddleware(e.InternalAuthToken)

	// Handler
	d.SampleHandler = site.NewSampleHandler(svc)
}

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
	"github.com/aikizoku/rabbitgo/appengine/api/src/usecase"
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
	cFirebaseAuth := firebaseauth.NewClient(e.ProjectID)
	cFirestore := cloudfirestore.NewClient(e.ProjectID)
	var cLog log.Writer
	if deploy.IsLocal() {
		cLog = log.NewWriterStdout()
	} else {
		cLog = log.NewWriterStackdriver(e.ProjectID)
	}
	cPubsub := cloudpubsub.NewClient(e.ProjectID, []string{
		images.ConverterTopicID,
	})
	cImages := images.NewClient(cPubsub)

	// Repository
	rSample := repository.NewSample(cFirestore, cImages)

	// Usecase
	uSample := usecase.NewSample(rSample)

	// Service
	var sFirebaseAuth firebaseauth.Service
	if deploy.IsProduction() {
		sFirebaseAuth = firebaseauth.NewService(cFirebaseAuth)
	} else {
		sFirebaseAuth = firebaseauth.NewServiceDebug(cFirebaseAuth, map[string]interface{}{})
	}
	sSample := service.NewSample(uSample, rSample)

	// Middleware
	d.Accesscontrol = accesscontrol.NewMiddleware(nil)
	d.Log = log.NewMiddleware(cLog, e.MinLogSeverity)
	d.FirebaseAuth = firebaseauth.NewMiddleware(sFirebaseAuth, false)
	d.InternalAuth = internalauth.NewMiddleware(e.InternalAuthToken)

	// Handler
	d.SampleHandler = site.NewSampleHandler(sSample)
}

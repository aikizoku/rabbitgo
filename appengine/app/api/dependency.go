package main

import (
	"github.com/aikizoku/beego/src/handler/api"
	"github.com/aikizoku/beego/src/lib/firebaseauth"
	"github.com/aikizoku/beego/src/lib/headerparams"
	"github.com/aikizoku/beego/src/lib/jsonrpc2"
	"github.com/aikizoku/beego/src/repository"
	"github.com/aikizoku/beego/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	DummyFirebaseAuth *firebaseauth.Middleware
	FirebaseAuth      *firebaseauth.Middleware
	DummyHeaderParams *headerparams.Middleware
	HeaderParams      *headerparams.Middleware
	JSONRPC2          *jsonrpc2.Middleware
	SampleHandler     *api.SampleHandler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject() {
	// Config
	// dbCfg := config.GetCSQLConfig("sample")

	// Lib
	// dbConn := cloudsql.NewCSQLClient(dbCfg)

	// Repository
	repo := repository.NewSample(nil)

	// Service
	dfaSvc := firebaseauth.NewDummyService()
	faSvc := firebaseauth.NewService()
	dhpSvc := headerparams.NewDummyService()
	hpSvc := headerparams.NewService()
	svc := service.NewSample(repo)

	// Middleware
	d.DummyFirebaseAuth = firebaseauth.NewMiddleware(dfaSvc)
	d.FirebaseAuth = firebaseauth.NewMiddleware(faSvc)
	d.DummyHeaderParams = headerparams.NewMiddleware(dhpSvc)
	d.HeaderParams = headerparams.NewMiddleware(hpSvc)

	// JSONRPC2
	d.JSONRPC2 = jsonrpc2.NewMiddleware()
	d.JSONRPC2.Register("sample", api.NewSampleJSONRPC2Handler(svc))

	// Handler
	d.SampleHandler = api.NewSampleHandler(svc)
}

package main

import (
	"github.com/aikizoku/beego/src/config"
	"github.com/aikizoku/beego/src/handler/api"
	"github.com/aikizoku/beego/src/infrastructure"
	"github.com/aikizoku/beego/src/lib/jsonrpc2"
	"github.com/aikizoku/beego/src/middleware"
	"github.com/aikizoku/beego/src/repository"
	"github.com/aikizoku/beego/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	FirebaseAuth          *middleware.FirebaseAuth
	JSONRPC2              *jsonrpc2.Middleware
	SampleHandler         *api.SampleHandler
	SampleJSONRPC2Handler *api.SampleJSONRPC2Handler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject() {
	// Config
	dbCfg := config.GetCSQLConfig("sample")

	// Lib
	dbConn := infrastructure.NewCSQLClient(dbCfg)

	// Repository
	repo := repository.NewSampleRepository(dbConn)

	// Service
	auth := service.NewAuthenticator()
	svc := service.NewSampleService(repo)

	// Middleware
	d.FirebaseAuth = middd
	d.JSONRPC2 = jsonrpc2.NewMiddleware()

	// Handler
	d.SampleHandler = &api.SampleHandler{
		Svc: svc,
	}
}

package main

import (
	"github.com/aikizoku/beego/src/handler/api"
	"github.com/aikizoku/beego/src/middleware"
	"github.com/aikizoku/beego/src/repository"
	"github.com/aikizoku/beego/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	FirebaseAuth  *middleware.FirebaseAuth
	SampleHandler *api.SampleHandler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject() {
	// Repository
	repo := repository.NewSampleRepository()

	// Service
	auth := service.NewAuthenticator()
	svc := service.NewSampleService(repo)

	// Middleware
	d.FirebaseAuth = &middleware.FirebaseAuth{Authenticator: auth}

	// Handler
	d.SampleHandler = &api.SampleHandler{Authenticator: auth, SampleService: svc}
}

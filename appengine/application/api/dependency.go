package main

import (
	"github.com/phi-jp/lightning-backend/src/handler/api"
	"github.com/phi-jp/lightning-backend/src/middleware"
	"github.com/phi-jp/lightning-backend/src/repository"
	"github.com/phi-jp/lightning-backend/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	FirebaseAuth *middleware.FirebaseAuth
	BeegoHandler *api.BeegoHandler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject() {
	// Repository
	bRepo := repository.NewBeego()

	// Service
	auth := service.NewAuthenticator()
	bSvc := service.NewBeego(bRepo)

	// Middleware
	d.FirebaseAuth = &middleware.FirebaseAuth{Authenticator: auth}

	// Handler
	d.BeegoHandler = &api.BeegoHandler{Authenticator: auth, BeegoService: bSvc}
}

package main

import (
	"github.com/phi-jp/lightning-backend/src/config"
	"github.com/phi-jp/lightning-backend/src/infrastructure"
	"github.com/phi-jp/lightning-backend/src/repository"
	"github.com/phi-jp/lightning-backend/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	BeegoAdminHandler *worker.BeegoAdminHandler
	BeegoTaskHandler  *worker.BeegoTaskHandler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject() {
	// Infrastructure
	req := infrastructure.NewHTTP(config.HTTPRequestTimeout)

	// Repository
	bRepo := repository.NewBeego()

	// Service
	auth := service.NewAuthenticator()
	bSvc := service.NewBeego(bRepo)

	// Handler
	d.BeegoAdminHandler = &worker.BeegoAdminHandler{BeegoService: bSvc}
	d.BeegoTaskHandler = &worker.BeegoTaskHandler{BeegoService: bSvc}
}

package main

import (
	"github.com/aikizoku/beego/src/config"
	"github.com/aikizoku/beego/src/handler/worker"
	"github.com/aikizoku/beego/src/infrastructure"
	"github.com/aikizoku/beego/src/repository"
	"github.com/aikizoku/beego/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	AdminHandler *worker.AdminHandler
	BeegoHandler *worker.BeegoHandler
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
	d.BeegoAdminHandler = &worker.AdminHandler{BeegoService: bSvc}
	d.BeegoTaskHandler = &worker.BeegoHandler{BeegoService: bSvc}
}

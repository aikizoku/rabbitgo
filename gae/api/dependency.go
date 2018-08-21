package main

import (
	"github.com/phi-jp/lightning-backend/src/handler/api"
	"github.com/phi-jp/lightning-backend/src/middleware"
	"github.com/phi-jp/lightning-backend/src/repository"
	"github.com/phi-jp/lightning-backend/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	FirebaseAuth    *middleware.FirebaseAuth
	LaunchHandler   *api.LaunchHandler
	ArticlesHandler *api.ArticlesHandler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject() {
	// Repository
	fdbRepo := repository.NewFeedDB()
	adbRepo := repository.NewArticleDB()
	ldbRepo := repository.NewArticleListDB()
	asRepo := repository.NewArticleSearch()
	ahdbRepo := repository.NewArticleHistoryDB()
	aldbRepo := repository.NewArticleLaterDB()

	// Service
	auth := service.NewAuthenticator()
	aSvc := service.NewArticle(fdbRepo, adbRepo, ahdbRepo, aldbRepo)
	sSvc := service.NewSearch(aSvc, asRepo)
	lSvc := service.NewList(aSvc, sSvc, adbRepo, ldbRepo, ahdbRepo, aldbRepo)

	// Handler
	d.FirebaseAuth = &middleware.FirebaseAuth{Authenticator: auth}
	d.LaunchHandler = &api.LaunchHandler{Authenticator: auth}
	d.ArticlesHandler = &api.ArticlesHandler{List: lSvc, Search: sSvc, Article: aSvc}
}

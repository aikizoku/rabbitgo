package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/phi-jp/lightning-backend/src/handler"
	"github.com/phi-jp/lightning-backend/src/middleware"
)

// Routing ... ハンドラをルーティングする
func Routing(r *chi.Mux, d *Dependency) {
	// アクセスコントロール
	r.Use(middleware.AccessControl)
	// Ping
	r.Get("/ping", handler.PingHandler)

	// ダミー認証
	r.Route("/internal/v1", func(r chi.Router) {
		r.Use(d.FirebaseAuth.DummyAuthentication)
		r.Use(middleware.GetDummyHeaderParams)
		subRouting(r, d)
	})

	// 認証
	r.Route("/v1", func(r chi.Router) {
		r.Use(d.FirebaseAuth.Authentication)
		r.Use(middleware.GetHeaderParams)
		subRouting(r, d)
	})

	http.Handle("/", r)
}

func subRouting(r chi.Router, d *Dependency) {
	r.Post("/launch", d.LaunchHandler.Launch)

	r.Route("/articles", func(r chi.Router) {
		r.Get("/list/{name}", d.ArticlesHandler.GetList)
		r.Get("/search", d.ArticlesHandler.GetSearch)
		r.Get("/relation", d.ArticlesHandler.GetRelationArticles)
	})

	r.Route("/temp", func(r chi.Router) {
		r.Get("/articles/list/{name}", d.ArticlesHandler.TempGetList)
	})

	r.Route("/user", func(r chi.Router) {
		r.Route("/articles", func(r chi.Router) {
			r.Post("/read/{article_id}", d.ArticlesHandler.PutRead)
			r.Get("/read", d.ArticlesHandler.GetHistories)
			r.Get("/later", d.ArticlesHandler.GetLaters)
			r.Post("/later/{article_id}", d.ArticlesHandler.PutLater)
			r.Delete("/later/{article_id}", d.ArticlesHandler.DeleteLater)
			r.Delete("/later/read", d.ArticlesHandler.DeleteLaterRead)
		})
	})
}

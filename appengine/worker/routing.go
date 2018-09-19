package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/phi-jp/lightning-backend/src/handler"
)

// Routing ... アプリのルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	r.Get("/ping", handler.PingHandler)

	r.Route("/admin", func(r chi.Router) {
		r.Get("/migration", d.AdminHandler.Migration)
		r.Get("/reindex", d.AdminHandler.Reindex)
		r.Get("/keywords", d.AdminHandler.ArticleKeywords)
	})

	r.Route("/task", func(r chi.Router) {
		r.Route("/fetch", func(r chi.Router) {
			r.Get("/feeds", d.FetchHandler.Feeds)
			r.Post("/feed", d.FetchHandler.Feed)
		})
		r.Route("/generate", func(r chi.Router) {
			r.Get("/article-lists", d.ListGenerateHandler.ArticleLists)
			r.Post("/latest-article-list", d.ListGenerateHandler.LatestArticleList)
		})
		r.Route("/index", func(r chi.Router) {
			r.Post("/article", d.IndexHandler.Article)
		})
		r.Route("/sweep", func(r chi.Router) {
			r.Get("/old-articles", d.SweepHandler.OldArticles)
		})
		r.Route("/analyze", func(r chi.Router) {
			r.Post("/article/keywords", d.AnalyzeHandler.ArticleKeywords)
		})
	})

	http.Handle("/", r)
}

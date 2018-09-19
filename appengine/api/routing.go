package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/phi-jp/lightning-backend/src/handler"
	"github.com/phi-jp/lightning-backend/src/middleware"
)

// Routing ... アプリのルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	// アクセスコントロール
	r.Use(middleware.AccessControl)

	// Ping
	r.Get("/ping", handler.PingHandler)

	// 認証なし
	r.Route("/internal/v1", func(r chi.Router) {
		r.Use(d.FirebaseAuth.DummyAuthentication)
		r.Use(middleware.GetDummyHeaderParams)
		subRouting(r, d)
	})

	// 認証あり
	r.Route("/v1", func(r chi.Router) {
		r.Use(d.FirebaseAuth.Authentication)
		r.Use(middleware.GetHeaderParams)
		subRouting(r, d)
	})

	http.Handle("/", r)
}

func subRouting(r chi.Router, d *Dependency) {
	r.Get("/beego", d.BeegoHandler.Beego)
	r.Post("/rpc", d.BeegoHandler)
}

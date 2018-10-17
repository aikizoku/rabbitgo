package main

import (
	"net/http"

	"github.com/aikizoku/beego/src/handler"
	"github.com/aikizoku/beego/src/middleware"
	"github.com/go-chi/chi"
)

// Routing ... ルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	// アクセスコントロール
	r.Use(middleware.AccessControl)

	// Ping
	r.Get("/ping", handler.Ping)

	// 認証なし
	r.Route("/internal/v1", func(r chi.Router) {
		r.Use(d.FirebaseAuth.DummyAuth)
		r.Use(middleware.GetDummyHeaderParams)
		subRouting(r, d)
	})

	// 認証あり
	r.Route("/v1", func(r chi.Router) {
		r.Use(d.FirebaseAuth.Auth)
		r.Use(middleware.GetHeaderParams)
		subRouting(r, d)
	})

	http.Handle("/", r)
}

func subRouting(r chi.Router, d *Dependency) {
	// API
	r.Get("/sample", d.SampleHandler.Sample)

	// API(JSONRPC2)
	r.Route("/rpc", func(r chi.Router) {
		r.Post("/", handler.Empty)
	})

}

package main

import (
	"net/http"

	"github.com/aikizoku/merlin/src/handler"
	"github.com/aikizoku/merlin/src/lib/log"
	"github.com/aikizoku/merlin/src/middleware"
	"github.com/go-chi/chi"
)

// Routing ... ルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	// アクセスコントロール
	r.Use(middleware.AccessControl)

	// ログ
	r.Use(log.Handle)

	r.Route("/v1", func(r chi.Router) {
		r.Use(d.FirebaseAuth.Handle)
		r.Use(d.HTTPHeader.Handle)
		subRouting(r, d)
	})

	// Ping
	r.Get("/ping", handler.Ping)

	http.Handle("/", r)
}

func subRouting(r chi.Router, d *Dependency) {
	// API
	r.Get("/sample", d.SampleHandler.Sample)

	// API(JSONRPC2)
	r.Post("/rpc", d.JSONRPC2Handler.Handle)
}

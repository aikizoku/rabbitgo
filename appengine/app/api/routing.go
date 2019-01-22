package main

import (
	"net/http"

	"github.com/aikizoku/skgo/src/config"
	"github.com/aikizoku/skgo/src/handler"
	"github.com/aikizoku/skgo/src/lib/log"
	"github.com/aikizoku/skgo/src/middleware"
	"github.com/go-chi/chi"
)

// Routing ... ルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	// アクセスコントロール
	r.Use(middleware.AccessControl)

	// ログ
	r.Use(log.Handle)

	// 認証なし(Stagingのみ)
	if config.IsEnvStaging() {
		r.Route("/noauth/v1", func(r chi.Router) {
			r.Use(d.DummyFirebaseAuth.Handle)
			r.Use(d.DummyHTTPHeader.Handle)
			subRouting(r, d)
		})
	}

	// 認証あり
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
	r.Route("/rpc", func(r chi.Router) {
		r.Use(d.JSONRPC2.Handle)
		r.Post("/", handler.Empty)
	})
}

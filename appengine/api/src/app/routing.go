package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rabee-inc/go-pkg/deploy"

	"github.com/aikizoku/rabbitgo/appengine/api/src/handler"
)

// Routing ... ルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	// ブラウザのCORS対応
	r.Use(d.Accesscontrol.Handle)

	// Log
	r.Use(d.Log.Handle)

	// Ping
	r.Get("/ping", handler.Ping)

	// 例: サブルーティング
	r.Route("/v1", func(r chi.Router) {
		// API
		r.Get("/sample", d.SampleHandler.Sample)
	})

	// 例: Stagingのみ適用したいルーティング
	if deploy.IsStaging() {
		// ここにルーティングを書く
	}

	http.Handle("/", r)
}

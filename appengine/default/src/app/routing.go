package app

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/aikizoku/rabbitgo/appengine/default/src/handler"
	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/accesscontrol"
	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/deploy"
)

// Routing ... ルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	// ブラウザのCORS対応
	r.Use(accesscontrol.Handle)

	// Log
	r.Use(d.Log.Handle)

	// Ping
	r.Get("/ping", handler.Ping)

	// 例: サブルーティング
	r.Route("/v1", func(r chi.Router) {
		// API
		r.Get("/sample", d.SampleHandler.Sample)

		// API(JSONRPC2)
		r.With(d.FirebaseAuth.Handle).Post("/rpc", d.JSONRPC2Handler.Handle)
	})

	// 例: Stagingのみ適用したいルーティング
	if deploy.IsStaging() {
		// ここにルーティングを書く
	}

	http.Handle("/", r)
}

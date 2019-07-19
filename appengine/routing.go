package main

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/aikizoku/rabbitgo/appengine/src/handler"
	"github.com/aikizoku/rabbitgo/appengine/src/lib/accesscontrol"
	"github.com/aikizoku/rabbitgo/appengine/src/lib/deploy"
)

// Routing ... ルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	// ブラウザのCORS対応
	r.Use(accesscontrol.Handle)

	// ログをリクエスト単位でまとめるため、情報をContextに保持する
	r.Use(d.Log.Handle)

	// 障害検知でサーバーの生存確認のため、pingリクエストを用意する
	r.Get("/ping", handler.Ping)

	// 例: サブルーティング
	r.Route("/v1", func(r chi.Router) {
		// API
		r.With(d.FirebaseAuth.Handle).Get("/sample", d.SampleHandler.Sample)

		// API(JSONRPC2)
		r.With(d.FirebaseAuth.Handle).Post("/rpc", d.JSONRPC2Handler.Handle)
	})

	// 例: Stagingのみ適用したいルーティング
	if deploy.IsStaging() {
		// ここにルーティングを書く
	}

	http.Handle("/", r)
}

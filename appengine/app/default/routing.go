package main

import (
	"net/http"

	"github.com/aikizoku/merlin/src/config"

	"github.com/aikizoku/merlin/src/handler"
	"github.com/aikizoku/merlin/src/lib/log"
	"github.com/aikizoku/merlin/src/middleware"
	"github.com/go-chi/chi"
)

// Routing ... ルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	/*
	 * ブラウザのCORS対応
	 */
	r.Use(middleware.AccessControl)

	/*
	 * ログをリクエスト単位でまとめるため、情報をContextに保持する
	 */
	r.Use(log.Handle)

	/*
	 * 障害検知でサーバーの生存確認のため、pingリクエストを用意する
	 */
	r.Get("/ping", handler.Ping)

	// 例: サブルーティング
	r.Route("/v1", func(r chi.Router) {
		/*
		 * FirebaseAuthentication
		 */
		r.Use(d.FirebaseAuth.Handle)
		/*
		 * HTTPHeaderから値を取得したい時はこちらを使う
		 */
		r.Use(d.HTTPHeader.Handle)

		// API
		r.Get("/sample", d.SampleHandler.Sample)

		// API(JSONRPC2)
		r.Post("/rpc", d.JSONRPC2Handler.Handle)
	})

	// 例: Stagingのみ適用したいルーティング
	if config.IsEnvStaging() {
		// ここにルーティングを書く
	}

	// 簡易テスト
	r.Get("/", handler.Empty)

	// 例: 個別にMiddlewareを適用したい場合
	r.With(d.FirebaseAuth.Handle).Get("/hoge", handler.Empty)

	http.Handle("/", r)
}

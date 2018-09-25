package main

import (
	"net/http"

	"github.com/aikizoku/beego/src/handler"
	"github.com/go-chi/chi"
)

// Routing ... アプリのルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	r.Get("/ping", handler.PingHandler)

	r.Route("/admin", func(r chi.Router) {
		r.Get("/migration", d.BeegoAdminHandler.Migration)
	})

	r.Route("/task", func(r chi.Router) {
		r.Route("/beego", func(r chi.Router) {
			r.Get("/beegos", d.BeegoTaskHandler.Beegos)
			r.Post("/beego", d.BeegoTaskHandler.Beego)
		})
	})

	http.Handle("/", r)
}

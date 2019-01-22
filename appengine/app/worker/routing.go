package main

import (
	"net/http"

	"github.com/aikizoku/skgo/src/handler"
	"github.com/go-chi/chi"
)

// Routing ... ルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	r.Get("/ping", handler.Ping)

	r.Route("/admin", func(r chi.Router) {
		r.Route("/migration", func(r chi.Router) {
			r.Get("/masterdata", d.AdminHandler.MigrateMasterData)
			r.Get("/testdata", d.AdminHandler.MigrateTestData)
		})
	})

	r.Route("/task", func(r chi.Router) {
		r.Route("/sample", func(r chi.Router) {
			r.Get("/cron", d.SampleHandler.Cron)
			r.Post("/taskqueue", d.SampleHandler.TaskQueue)
		})
	})

	http.Handle("/", r)
}

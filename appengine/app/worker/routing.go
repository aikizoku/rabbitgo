package main

import (
	"net/http"

	"github.com/aikizoku/beego/src/handler"
	"github.com/go-chi/chi"
)

// Routing ... ルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	r.Get("/ping", handler.PingHandler)

	r.Route("/admin", func(r chi.Router) {
		r.Route("/migration", func(r chi.Router) {
			r.Get("/masterdata", d.AdminHandler.MigrateMasterData)
			r.Get("/testdata", d.AdminHandler.MigrateTestData)
		})
	})

	r.Route("/task", func(r chi.Router) {
		r.Route("/sample", func(r chi.Router) {
			r.Get("/cron", d.SampleHandler.CronHandler)
			r.Post("/taskqueue", d.SampleHandler.TaskQueueHandler)
		})
	})

	http.Handle("/", r)
}

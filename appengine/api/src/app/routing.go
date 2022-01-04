package app

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/aikizoku/rabbitgo/appengine/api/src/handler"
)

// Routing ... ルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	r.Use(d.Accesscontrol.Handle)
	r.Use(d.Log.Handle)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/sample", d.SampleHandler.Sample)
	})

	r.Get("/ping", handler.Ping)

	http.Handle("/", r)
}

package web

import "github.com/go-chi/chi"

func (h *Sample) Routing(r *chi.Mux) {

	r.Get("/", h.Hoge)
}

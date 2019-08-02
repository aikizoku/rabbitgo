package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	// Environment
	e := &Environment{}
	e.Get()

	// Dependency
	d := &Dependency{}
	d.Inject(e)

	// Routing
	r := chi.NewRouter()
	Routing(r, d)

	// Run
	http.ListenAndServe(fmt.Sprintf(":%d", 8080), r)
}

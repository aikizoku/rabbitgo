package main

import (
	"github.com/go-chi/chi"
	"google.golang.org/appengine"
)

func main() {
	// Dependency
	d := &Dependency{}
	d.Inject()

	// Routing
	r := chi.NewRouter()
	Routing(r, d)

	// Run
	appengine.Main()
}

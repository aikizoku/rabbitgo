package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/aikizoku/rabbitgo/appengine/default/src/app"
)

func main() {
	// Environment
	e := &app.Environment{}
	e.Get()

	// Dependency
	d := &app.Dependency{}
	d.Inject(e)

	// Routing
	r := chi.NewRouter()
	app.Routing(r, d)

	// Run
	http.ListenAndServe(fmt.Sprintf(":%d", e.Port), r)
}

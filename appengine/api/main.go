package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/aikizoku/rabbitgo/appengine/api/src/app"
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
	if err := http.ListenAndServe(fmt.Sprintf(":%d", e.Port), r); err != nil {
		log.Fatal(err)
	}
}

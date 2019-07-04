module github.com/aikizoku/merlin/appengine/app/worker

replace github.com/aikizoku/merlin/src => ../../../src

require (
	github.com/aikizoku/merlin/src v1.0.0
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/unrolled/render v1.0.0 // indirect
	google.golang.org/appengine v1.6.1
)

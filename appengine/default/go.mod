module github.com/aikizoku/rabbitgo/appengine/default

go 1.12

replace github.com/aikizoku/rabbitgo/appengine/src => ../src

require (
	cloud.google.com/go v0.43.0
	firebase.google.com/go v3.8.1+incompatible
	github.com/aikizoku/rabbitgo/appengine/src v1.0.0
	github.com/davecgh/go-spew v1.1.1
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/joho/godotenv v1.3.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/rs/xid v1.2.1
	github.com/unrolled/render v1.0.0
	golang.org/x/text v0.3.2
	google.golang.org/api v0.7.0
)

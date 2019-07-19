module github.com/aikizoku/rabbitgo/appengine/default

go 1.12

replace github.com/aikizoku/rabbitgo/appengine/src => ../src

require (
	cloud.google.com/go v0.41.0 // indirect
	firebase.google.com/go v3.8.1+incompatible // indirect
	github.com/aikizoku/rabbitgo/appengine/src v1.0.0
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/joho/godotenv v1.3.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/rs/xid v1.2.1 // indirect
	github.com/unrolled/render v1.0.0 // indirect
	github.com/vvakame/sdlog v0.0.0-20190523062053-be70263e9c6c // indirect
)

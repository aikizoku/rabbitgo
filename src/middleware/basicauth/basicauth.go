package basicauth

import "net/http"

var accounts = map[string]string{
	"hoge": "hoge",
}

func basicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", "Basic")
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "basic auth required", http.StatusUnauthorized)
			return
		}
		if accounts[user] != password {
			http.Error(w, "failed to auth", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

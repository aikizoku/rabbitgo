package middleware

import "net/http"

var accounts = map[string]string{
	"hoge": "hoge",
}

// BasicAuth ... ベーシック認証機能を提供するミドルウェア
func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", "Basic")
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "basic auth required.", http.StatusUnauthorized)
			return
		}
		if accounts[user] != password {
			http.Error(w, "basic auth error.", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

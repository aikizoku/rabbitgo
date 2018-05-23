package middleware

import "net/http"

// BasicAuthenticator ... ベーシック認証機能を提供するミドルウェア
type BasicAuthenticator struct {
	Accounts map[string]string
}

// Handle ... ハンドラ
func (a *BasicAuthenticator) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", "Basic")
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "basic auth required.", http.StatusUnauthorized)
			return
		}
		if a.Accounts[user] != password {
			http.Error(w, "basic auth error.", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// NewBasicAuthenticator ... BasicAuthenticatorを作成する
func NewBasicAuthenticator(accounts map[string]string) *BasicAuthenticator {
	return &BasicAuthenticator{
		Accounts: accounts,
	}
}

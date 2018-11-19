package internalauth

import "net/http"

// Middleware ... 内部認証機能を提供するミドルウェア
type Middleware struct {
	AuthKey string
}

// Handle ... ハンドラ
func (m *Middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ah := r.Header.Get("Authorization")
		if ah == "" || ah != m.AuthKey {
			http.Error(w, "internal auth error.", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// NewMiddleware ... Middlewareを作成する
func NewMiddleware(authKey string) *Middleware {
	return &Middleware{
		AuthKey: authKey,
	}
}

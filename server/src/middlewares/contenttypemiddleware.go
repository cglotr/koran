package middlewares

import "net/http"

// ContentTypeMiddleware .
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json;charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

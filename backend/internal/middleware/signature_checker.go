package middleware

import (
	"net/http"
)

func SignatureChecker(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		//TODO: write logic here
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}

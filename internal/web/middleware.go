package web

import (
	"net/http"

	"go.uber.org/zap"
)

func WithLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		zap.L().Info(
			"Request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Any("query", r.URL.Query()),
			zap.Int64("body", r.ContentLength),
		)
		next.ServeHTTP(w, r)
	})
}

func WithContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

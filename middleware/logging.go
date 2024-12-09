package middleware

import (
	"log"
	"net/http"
	"time"
)

type WrappedWritter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *WrappedWritter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &WrappedWritter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		log.Println(r.Method, r.URL.Path, wrapped.StatusCode, time.Since(start))
	})
}

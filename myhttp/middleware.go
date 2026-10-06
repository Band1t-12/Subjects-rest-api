package myhttp

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		fmt.Println(r.Method, r.URL.Path, time.Since(start))
	})
}

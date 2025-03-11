// common/middleware/logging.go
package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the details of incoming requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Method: %s, Path: %s, Duration: %s\n", r.Method, r.URL.Path, time.Since(start))
	})
}

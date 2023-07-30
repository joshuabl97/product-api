package handlers

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

// HandlerLogger logs information about incoming requests
func HandlerLogger(next http.Handler, l *zerolog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Logging the request details
		loggerMiddlewareMiddleware(start, l, r)
	})
}

// LoggingMiddleware accepts a HandlerFunc and a *zerolog.Logger,
// and returns a new HandlerFunc that logs information about incoming requests.
func LoggingMiddleware(handler http.HandlerFunc, l *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the provided handler function
		handler(w, r)

		// Logging the request details
		loggerMiddlewareMiddleware(start, l, r)
	}
}

func loggerMiddlewareMiddleware(start time.Time, l *zerolog.Logger, r *http.Request) {
	// Logging the request details
	l.Info().
		Str("method", r.Method).
		Str("path", r.URL.Path).
		Str("remote_addr", r.RemoteAddr).
		Dur("duration", time.Since(start)).
		Msg("request handled")
}

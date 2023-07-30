package handlers

import (
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

// HandlerLogger logs information about incoming requests
func HandlerLogger(next http.Handler, logger *zerolog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Logging the request details
		logger.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Str("remote_addr", r.RemoteAddr).
			Dur("duration", time.Since(start)).
			Msg("request handled")
	})
}

// LoggingMiddleware accepts a HandlerFunc and a *zerolog.Logger,
// and returns a new HandlerFunc that logs information about incoming requests.
func LoggingMiddleware(handler http.HandlerFunc, logger *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the provided handler function
		handler(w, r)

		// Logging the request details
		logger.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Str("remote_addr", r.RemoteAddr).
			Dur("duration", time.Since(start)).
			Msg("request handled")
	}
}

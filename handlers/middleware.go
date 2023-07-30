package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/joshuabl97/product-api/data"
	"github.com/rs/zerolog"
)

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Error().
				Err(err).
				Msg("Unable to decode JSON")

			http.Error(
				rw,
				"Unable to decode json",
				http.StatusBadRequest)
		}

		// validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Error().
				Err(err).
				Msg("Unable to validate product")

			http.Error(
				rw,
				fmt.Sprintf("Error validitating product: %s\n", err),
				http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}

// LoggingMiddleware logs information about incoming requests
func LoggingMiddleware(next http.Handler, logger *zerolog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Logging the request details
		logger.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Str("remote_addr", r.RemoteAddr).
			Dur("duration", time.Since(start)).
			Msg("request handled")

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

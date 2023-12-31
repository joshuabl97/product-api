package handlers

import (
	"context"
	"fmt"
	"github.com/joshuabl97/product-api/data"
	"net/http"
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

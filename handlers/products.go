package handlers

import (
	"github.com/joshuabl97/product-api/data"
	"github.com/rs/zerolog"
)

// some message
// swagger:response productsResponse
type productsResponse struct {
	// all products in the data store
	// in: body
	Body []data.Product
}

// Products is a http.Handler
type Products struct {
	l *zerolog.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *zerolog.Logger) *Products {
	return &Products{l}
}

type KeyProduct struct{}

package handlers

import (
	"github.com/rs/zerolog"
)

// Products is a http.Handler
type Products struct {
	l *zerolog.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *zerolog.Logger) *Products {
	return &Products{l}
}

type KeyProduct struct{}

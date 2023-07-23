package handlers

import (
	"fmt"
	"net/http"

	"github.com/joshuabl97/product-api/data"
)

// add products adds a product to the productsList
// A product must contain a valid name, price, and SKU
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Info().Msg("Handle POST Products")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Info().
		Str("product", fmt.Sprintf("%v", prod)).
		Msg("Product successfully validated")

	data.AddProduct(prod)
	//data.AddProduct(prod)
}

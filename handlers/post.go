package handlers

import (
	"net/http"

	"github.com/joshuabl97/product-api/data"
)

// add products adds a product to the productsList
// A product must contain a valid name, price, and SKU
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("Prod: %#v\n", prod)
	data.AddProduct(prod)
	//data.AddProduct(prod)
}

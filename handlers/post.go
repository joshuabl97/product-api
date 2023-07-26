package handlers

import (
	"fmt"
	"net/http"

	"github.com/joshuabl97/product-api/data"
)

// swagger:route POST /products products addProduct
// Adds a product to the system
//     Parameters:
//       + name: name
//         in: query
//         description: name of product
//         required: true
//         type: string
//         format: string
//       + name: description
//         in: query
//         description: product description
//         required: false
//         type: string
//         format: string
//       + name: price
//         in: query
//         description: product description
//         required: true
//         type: float
//         format: float32
//       + name: sku
//         in: query
//         description: product sku
//         required: true
// 		   pattern: [a-z]+-[a-z]+-[a-z]+
//         type: string
//         format: string
// responses:
//	200: noContent

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

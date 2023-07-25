package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joshuabl97/product-api/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Info().Msg("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ProdsToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

// GetProduct returns a single product from the data store
func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Info().Str("id", mux.Vars(r)["id"]).Msg("Handle GET Product")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	prod, _, err := data.FindProduct(id)
	if err != nil {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
	}

	err = prod.ProdToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

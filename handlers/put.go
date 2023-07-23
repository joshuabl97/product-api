package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joshuabl97/product-api/data"
)

// handles PUT request to update a product
// PUT requests
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Info().Str("id", mux.Vars(r)["id"]).Msg("Handle PUT Products/{id}")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
	}

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	err = data.UpdateProduct(id, prod)
	if err != nil {
		http.Error(rw, "Unable to find product", http.StatusNotFound)
		return
	}

}

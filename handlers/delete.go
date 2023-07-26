package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joshuabl97/product-api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product from the data store
// responses:
//	200: noContent
// 400: noContent
// 404: noContent

// DeleteProduct deletes a product from the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Info().Str("id", mux.Vars(r)["id"]).Msg("Handle DELETE Products")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	if err := data.DeleteProduct(id); err != nil {
		p.l.Error().
			Err(err).
			Msg("Unable to delete product")

		http.Error(rw, "Unable to find/delete product", http.StatusNotFound)
		return
	}
}

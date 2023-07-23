package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joshuabl97/product-api/data"
)

func (p *Products) DeleteProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	p.l.Info().Str("id", string(id)).Msg("Handle DELETE Products")

	if err := data.DeleteProduct(id); err != nil {
		p.l.Error().
			Err(err).
			Msg("Unable to delete product")

		http.Error(rw, "Unable to find/delete product", http.StatusNotFound)
		return
	}
}

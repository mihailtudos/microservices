package handlers

import (
	"github.com/mihailtudos/microservices/data"
	"net/http"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns a list of products
//
// responses:
//		201: noContent
//		404: errorResponse
//		501: errorResponse

// DeleteProduct deletes a product from the DB
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		w.WriteHeader(http.StatusNotFound)
		_ = data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		w.WriteHeader(http.StatusInternalServerError)
		_ = data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

package handlers

import (
	"github.com/mihailtudos/microservices/data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//		200: productsResponse

// ListAll returns the products from the data store
func (p *Products) ListAll(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// set content-type header
	w.Header().Add("Content-Type", "application/json")

	// serialize the products list to JSON
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products listSingle
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusNotFound)
		_ = data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusInternalServerError)
		_ = data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	// set content-type header
	w.Header().Add("Content-Type", "application/json")
	err = data.ToJSON(prod, w)
	if err != nil {
		// we should never be here but log the error
		p.l.Println("[ERROR] serializing product", err)
	}
}

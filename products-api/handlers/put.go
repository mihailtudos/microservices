package handlers

import (
	data2 "github.com/mihailtudos/microservices/data"
	"net/http"
)

// swagger:route PUT /products/{id} products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {

	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data2.Product)
	p.l.Println("[DEBUG] updating record id", prod.ID)

	err := data2.UpdateProduct(prod)
	if err == data2.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)

		rw.WriteHeader(http.StatusNotFound)
		data2.ToJSON(&GenericError{Message: "Product not found in database"}, rw)
		return
	}

	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}

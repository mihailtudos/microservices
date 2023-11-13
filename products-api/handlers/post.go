package handlers

import (
	"github.com/mihailtudos/microservices/data"
	"net/http"
)

//swagger:route POST /products products createProduct
// Create new product
//
// responses:
//		201: productResponse
//		422: errorValidation
//		501: errorResponse

// Create updates a product to the received body
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("[DEBUG] Inserting product: %#v\n", prod)
	data.AddProduct(prod)
}

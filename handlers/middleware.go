package handlers

import (
	"context"
	"fmt"
	"github.com/mihailtudos/microservices/data"
	"net/http"
)

// MiddlewareProductValidation validates the product in the request and calls next if ok
func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product ", err)
			http.Error(w, "Error reading product", http.StatusBadRequest)
			return
		}

		// validate product
		ve := p.v.Validate(prod)
		if ve != nil {
			p.l.Println("[ERROR] validating product ", err)
			http.Error(w, fmt.Sprintf("Error validating product: %s", err.Error()), http.StatusUnprocessableEntity)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		rwc := r.WithContext(ctx)

		//Call the next handler, which can be another middleware in the chain,
		// or the final handler.
		next.ServeHTTP(w, rwc)
	})
}

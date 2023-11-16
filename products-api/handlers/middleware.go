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

func (p *Products) CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		origin := r.Header.Get("Origin")
		if isValidOrigin(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		}

		// Allow OPTIONS method without further processing
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func isValidOrigin(origin string) bool {
	// Check if the origin is in the list of allowed origins
	allowedOrigins := []string{"http://localhost:8081", "http://localhost:5173"}
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			return true
		}
	}
	return false
}

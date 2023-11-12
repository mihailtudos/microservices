package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mihailtudos/microservices/data"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GetProducts returns the products from the data store
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
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

// UpdateProduct updates a product to the received body
func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Id provided", http.StatusBadRequest)
	}

	p.l.Println("Handle PUT Product ", id)

	prod := &data.Product{}
	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}

// AddProduct updates a product to the received body
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
	w.WriteHeader(http.StatusCreated)
}

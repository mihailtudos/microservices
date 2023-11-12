package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/mihailtudos/microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	} else if r.Method == http.MethodPut {
		p.updateProduct(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	w.Header().Add("Content-Type", "application/json")

	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) updateProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("\n\n%#v\n\n", strings.Split(r.URL.Path, "/products/"))
}

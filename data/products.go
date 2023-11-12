package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a collection of Product
type Products []*Product

// ToJSON serializes the contents of the collection to JSON
func (p *Products) ToJSON(w io.Writer) error {
	// NewEncoder provides better performance than json.Marshal as it does not
	// have to buffer the output into a slice of bytes.
	// This reduces allocations and the overheads of the service
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

// productList is a Hard coded list of products for this data structure
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.55,
		SKU:         "LA1321",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "ES1123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

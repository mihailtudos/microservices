package main

import (
	"fmt"
	"github.com/mihailtudos/microservices/sdk/client"
	"github.com/mihailtudos/microservices/sdk/client/products"
	"testing"
)

func TestOutClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:8080")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()

	prod, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}

	for _, p := range prod.Payload {
		fmt.Printf("%#v\n", *p)
	}
}

package data

import "testing"

func TestCheckValidator(t *testing.T) {
	p := &Product{Name: "Nik", Price: 0.212, SKU: "aaa-aaa-aaaa"}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

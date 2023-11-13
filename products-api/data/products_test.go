package data

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckValidator(t *testing.T) {
	p := &Product{Name: "Nik", Price: 0.212, SKU: "aaa-aaa-aaaa"}

	v := NewValidation()
	err := v.Validate(p)
	if err != nil {
		t.Fatal(err)
	}
}

//func TestProductMissingPriceReturnsErr(t *testing.T) {
//	p := Product{
//		Name:  "abc",
//		Price: -1,
//	}
//
//	v := NewValidation()
//	err := v.Validate(p)
//	assert.Len(t, err, 1)
//}

func TestProductInvalidSKUReturnsErr(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: 1.22,
		SKU:   "abc",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

//func TestValidProductDoesNOTReturnsErr(t *testing.T) {
//	p := Product{
//		Name:  "abc",
//		Price: 1.22,
//		SKU:   "abc-efg-hji",
//	}
//
//	v := NewValidation()
//	err := v.Validate(p)
//	assert.Len(t, err, 1)
//}

func TestProductsToJSON(t *testing.T) {
	ps := []*Product{
		&Product{
			Name: "abc",
		},
	}

	b := bytes.NewBufferString("")
	err := ToJSON(ps, b)
	assert.NoError(t, err)
}

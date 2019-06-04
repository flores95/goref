package repo

import (
	"reflect"
	"testing"

	"github.com/flores95/goref/services/products/product"
)

func Test_Insert(t *testing.T) {
	r := Repo{}
	tests := []struct {
		name string
		p    product.Product
	}{
		{
			name: "Can insert a product",
			p: product.Product{UPC: "some-upc-to-test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.Insert(tt.p)
		})
	}
}

func TestRepo_GetAll(t *testing.T) {
	r := Repo{}
	r.Insert(product.Product{UPC: "some-upc-to-test-01"})
	r.Insert(product.Product{UPC: "some-upc-to-test-02"})
	r.Insert(product.Product{UPC: "some-upc-to-test-03"})
	tests := []struct {
		name string
		r    *Repo
		want []product.Product
	}{
		{
			name: "Get all products",
			r:    &r,
			want: []product.Product{
				product.Product{UPC: "some-upc-to-test-01"},
				product.Product{UPC: "some-upc-to-test-02"},
				product.Product{UPC: "some-upc-to-test-03"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repo.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

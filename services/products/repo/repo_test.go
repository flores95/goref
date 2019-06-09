package repo

import (
	"reflect"
	"testing"

	"github.com/flores95/goref/services/products/models"
)

func Test_Insert(t *testing.T) {
	r := Repo{}
	tests := []struct {
		name string
		p    models.Product
	}{
		{
			name: "Can insert a product",
			p: models.Product{UPC: "some-upc-to-test"},
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
	r.Insert(models.Product{UPC: "some-upc-to-test-01"})
	r.Insert(models.Product{UPC: "some-upc-to-test-02"})
	r.Insert(models.Product{UPC: "some-upc-to-test-03"})
	tests := []struct {
		name string
		r    *Repo
		want []models.Product
	}{
		{
			name: "Get all products",
			r:    &r,
			want: []models.Product{
				models.Product{UPC: "some-upc-to-test-01"},
				models.Product{UPC: "some-upc-to-test-02"},
				models.Product{UPC: "some-upc-to-test-03"},
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

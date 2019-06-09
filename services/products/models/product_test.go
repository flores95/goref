package models


import (
	"testing"
)

func Test_ProdccutHasUPC(t *testing.T) {
	tests := []struct {
		name string
		p    Product
		want string
	}{
		{
			name: "Product should have a UPC",
			p:    Product{UPC: "some-upc-number"},
			want: "some-upc-number",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.UPC; got != tt.want {
				t.Errorf("%v:: %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
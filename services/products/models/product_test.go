package models

import (
	"testing"

	"github.com/flores95/goref/frameworks"
)

func Test_ProductHasID(t *testing.T) {
	tests := []struct {
		name string
		p    *Product
		want frameworks.ID
	}{
		{
			name: "Product should have an ID",
			p:    NewProduct("test-id", "test-name", "test-description"),
			want: "test-id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.ID(); got != tt.want {
				t.Errorf("%v:: %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

package controllers

import (
	"reflect"
	"testing"

	"github.com/flores95/goref/frameworks/config"
)

func TestNewCLIProductController(t *testing.T) {
	cfg := config.NewBaseConfigurator("")
	tests := []struct {
		name string
		kvs  config.KVS
		want string
	}{
		{
			name: "Product Controller is created with a default product service endpoint",
			kvs:  config.KVS{},
			want: "http://localhost:4181/products",
		},
		{
			name: "Product Controller is created with an endpoint set in configuration data",
			kvs:  config.KVS{"PRODUCTS_EP": "http://configured/products"},
			want: "http://configured/products",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg.Load(tt.kvs)
			pc := NewCLIProductController(cfg)
			if got := pc.endpoint; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v :: got = %v :: want %v", tt.name, got, tt.want)
			}
		})
	}
}

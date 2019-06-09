package controllers

import (
	"reflect"
	"testing"

	"github.com/flores95/goref/frameworks/config"
)

func TestNewOrderController(t *testing.T) {
	cfg := config.NewBaseConfigurator("")
	tests := []struct {
		name string
		kvs  config.KVS
		want string
	}{
		{
			name: "Order Controller is created with a default order service endpoint",
			kvs:  config.KVS{},
			want: "http://localhost:4182/orders",
		},
		{
			name: "Order Controller is created with an endpoint set in configuration data",
			kvs:  config.KVS{"ORDERS_EP": "http://configured/orders"},
			want: "http://configured/orders",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg.Load(tt.kvs)
			pc := NewOrderController(cfg)
			if got := pc.endpoint; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v :: got = %v :: want %v", tt.name, got, tt.want)
			}
		})
	}
}

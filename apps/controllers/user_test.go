package controllers

import (
	"reflect"
	"testing"

	"github.com/flores95/goref/frameworks/config"
)

func TestNewUserController(t *testing.T) {
	cfg := config.NewBaseConfigurator("")
	tests := []struct {
		name string
		kvs  config.KVS
		want string
	}{
		{
			name: "User Controller is created with a default user service endpoint",
			kvs:  config.KVS{},
			want: "http://localhost:4180/users",
		},
		{
			name: "User Controller is created with an endpoint set in configuration data",
			kvs:  config.KVS{"USERS_EP": "http://configured/users"},
			want: "http://configured/users",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg.Load(tt.kvs)
			pc := NewUserController(cfg)
			if got := pc.endpoint; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v :: got = %v :: want %v", tt.name, got, tt.want)
			}
		})
	}
}

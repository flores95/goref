package config

import (
	"reflect"
	"testing"
)

func TestNewDotenvConfigurator(t *testing.T) {
	tests := []struct {
		name string
		want reflect.Type
	}{
		{name: "Can create a new DotenvConfigurator", want: reflect.TypeOf(new(DotenvConfigurator))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewDotenvConfigurator("")
			if got := reflect.TypeOf(c); got != tt.want {
				t.Errorf("%v :: got = %v :: want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestDotenvConfigurator_GetNamespace(t *testing.T) {
	tests := []struct {
		name      string
		namespace string
		want      string
	}{
		{name: "Configurator should set namespace in it's constructor", namespace: "TESTNS", want: "TESTNS"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := NewDotenvConfigurator(tt.namespace)
			if got := cfg.GetNamespace(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v :: got = %v :: want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestDotenvConfigurator_GetValue(t *testing.T) {
	conf := NewDotenvConfigurator("")
	conf.Load(KVS{"testkey": "testkey-value"})
	tests := []struct {
		name string
		c    Configurator
		key  string
		want string
	}{
		{
			name: "Given a key return it's value",
			c:    conf,
			key:  "testkey",
			want: "testkey-value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetValue(tt.key); got != tt.want {
				t.Errorf("%v :: got = %v :: want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestDotenvConfigurator_Load(t *testing.T) {
	conf := NewDotenvConfigurator("")
	kvs := KVS{"testkey": "testkey-value"}
	tests := []struct {
		name string
		c    Configurator
		kvs  KVS
	}{
		{name: "can load a key value map", c: conf, kvs: kvs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Load(tt.kvs)
			for key, wantValue := range kvs {
				gotValue := tt.c.GetValue(key)
				if wantValue != tt.c.GetValue(key) {
					t.Errorf("DotenvConfigurator.Load() = %v, want %v", gotValue, wantValue)
				}
			}
		})
	}
}

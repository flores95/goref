package config

import (
	"reflect"
	"testing"
)

func TestNewBaseConfigurator(t *testing.T) {
	tests := []struct {
		name string
		want Configurator
	}{
		{name: "Can create a new BaseConfigurator", want: new(BaseConfigurator)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBaseConfigurator(""); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v :: got = %v :: want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestBaseConfigurator_GetNamespace(t *testing.T) {
	tests := []struct {
		name      string
		namespace string
		want      string
	}{
		{name: "Configurator should set namespace in it's constructor", namespace: "TESTNS", want: "TESTNS"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := NewBaseConfigurator(tt.namespace)
			if got := cfg.GetNamespace(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v :: got = %v :: want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestBaseConfigurator_GetValue(t *testing.T) {
	cfg := NewBaseConfigurator("NS_")
	tests := []struct {
		name string
		kvs  KVS
		key  string
		want string
	}{
		{
			name: "Given a key return it's value",
			kvs:  KVS{"testkey": "testkey-value"},
			key:  "testkey",
			want: "testkey-value",
		},
		{
			name: "Return the value of a key without specifying a namespace",
			kvs:  KVS{"NS_testkey": "testkey-value"},
			key:  "testkey",
			want: "testkey-value",
		},
		{
			name: "Return the value of a namespaced key over a non-namespced key",
			kvs:  KVS{"NS_testkey": "testkey-value", "testkey": "wrong-value"},
			key:  "testkey",
			want: "testkey-value",
		},
		{
			name: "Return a non namespace value of a key if one with namespace is not present",
			kvs:  KVS{"NS_wrongkey": "wrong-value", "testkey": "testkey-value"},
			key:  "testkey",
			want: "testkey-value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg.Load(tt.kvs)
			if got := cfg.GetValue(tt.key); got != tt.want {
				t.Errorf("%v :: got = %v :: want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestBaseConfigurator_Load(t *testing.T) {
	conf := NewBaseConfigurator("")
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
			for key, want := range kvs {
				got := tt.c.GetValue(key)
				if want != tt.c.GetValue(key) {
					t.Errorf("%v :: got = %v :: want %v", tt.name, got, want)
				}
			}
		})
	}
}

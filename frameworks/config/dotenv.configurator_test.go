package config

import (
	"reflect"
	"testing"
)

func TestNewDotenvConfigurator(t *testing.T) {
	wanted := DotenvConfigurator{kvs: map[string]string{}}
	tests := []struct {
		name string
		want Configurator
	}{
		{name: "Can create a new DotenvConfigurator", want: &wanted},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDotenvConfigurator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDotenvConfigurator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDotenvConfigurator_GetValue(t *testing.T) {
	conf := NewDotenvConfigurator()
	conf.Load(map[string]string{"testkey": "testkey-value"})
	tests := []struct {
		name      string
		c         Configurator
		key       string
		wantValue string
	}{
		{
			name:      "Given a key return it's value",
			c:         conf,
			key:       "testkey",
			wantValue: "testkey-value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.c.GetValue(tt.key); gotValue != tt.wantValue {
				t.Errorf("DotenvConfigurator.GetValue() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestDotenvConfigurator_Load(t *testing.T) {
	conf := NewDotenvConfigurator()
	kvs := map[string]string{"testkey": "testkey-value"}
	tests := []struct {
		name string
		c    Configurator
		kvs  map[string]string
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

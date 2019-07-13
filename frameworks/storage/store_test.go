package storage

import (
	"reflect"
	"testing"

	"github.com/flores95/goref/frameworks/config"
)

func TestNewStore(t *testing.T) {
	conf := config.NewBaseConfigurator("")
	tests := []struct {
		name string
		kvs  config.KVS
		want Store
	}{
		{
			name: "should create a new memory store with default settings with no config data",
			kvs:  config.KVS{},
			want: NewMemoryStore("", "", 0),
		},
		{
			name: "should create a new memory store with specified config settings",
			kvs:  config.KVS{"STORE_NAME": "TEST_MEMORY_STORE", "STORE_ID_PREFIX": "TESTID", "STORE_ID_SEED": "0"},
			want: NewMemoryStore("TEST_MEMORY_STORE", "TESTID", 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf.Load(tt.kvs)
			if got := NewStore(conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

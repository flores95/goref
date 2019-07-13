package users

import (
	"reflect"
	"testing"

	"github.com/flores95/goref/frameworks/config"
	"github.com/flores95/goref/frameworks/log"
	"github.com/flores95/goref/frameworks/storage"
)

func TestNewService(t *testing.T) {
	conf := config.NewBaseConfigurator("")

	tests := []struct {
		name string
		kvs  config.KVS
	}{
		{
			name: "should create default logger and store with no configuration settings",
			kvs:  config.KVS{},
		},
		{
			name: "should create the configured logger and store given the right configuration settings",
			kvs: config.KVS{
				"STORE_USING":     "MEMORY",
				"STORE_NAME":      "TEST STORAGE IN MEMORY",
				"STORE_ID_PREFIX": "ID",
				"STORE_ID_SEED":   "2001",
				"LOG_USING":       "CONSOLE",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf.Load(tt.kvs)
			wantLogger := log.NewLogger(conf)
			wantStore := storage.NewStore(conf)
			svc := NewService(conf)
			if gotLogger := svc.log; !reflect.DeepEqual(gotLogger, wantLogger) {
				t.Errorf("%v :: got: %v :: want: %v", tt.name, gotLogger, wantLogger)
			}
			if gotStore := svc.store; !reflect.DeepEqual(gotStore, wantStore) {
				t.Errorf("%v :: got: %v :: want: %v", tt.name, gotStore, wantStore)
			}
		})
	}
}

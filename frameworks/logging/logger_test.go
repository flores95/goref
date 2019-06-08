package logging

import (
	"reflect"
	"testing"

	"github.com/flores95/goref/frameworks/config"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name string
		cfg  config.Configurator
		kvs  config.KVS
		want reflect.Type
	}{
		{
			name: "Logger Factory provides console logger by default",
			kvs:  config.KVS{},
			cfg:  new(config.BaseConfigurator),
			want: reflect.TypeOf(new(ConsoleLogger)),
		},
		{
			name: "Logger Factory can create a console logger",
			kvs:  config.KVS{"LOG_USING": "CONSOLE"},
			cfg:  new(config.BaseConfigurator),
			want: reflect.TypeOf(new(ConsoleLogger)),
		},
		{
			name: "Logger Factory can create a file logger",
			kvs:  config.KVS{"LOG_USING": "file"},
			cfg:  new(config.BaseConfigurator),
			want: reflect.TypeOf(new(FileLogger)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cfg.Load(tt.kvs)
			l := NewLogger(tt.cfg)
			if got := reflect.TypeOf(l); tt.want != got {
				t.Errorf("%v :: got = %v :: want %v", tt.name, got, tt.want)
			}
		})
	}
}

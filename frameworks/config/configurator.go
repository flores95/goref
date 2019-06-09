package config

// Configurator is an interface used to define the basic requirements of a configurator
type Configurator interface {
	GetValue(string) string
	GetNamespace() string
	Load(KVS)
}

// Configurable can be used to define objects that use a configurator
type Configurable interface {
	Configure(Configurator)
}

// KVS is a simple string based key value map
type KVS map[string]string

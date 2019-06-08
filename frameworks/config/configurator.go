package config

type Configurator interface {
	GetValue(string) string
	GetNamespace() string
	Load(KVS)
}

type Configurable interface {
	Configure(Configurator)
}

type KVS map[string]string

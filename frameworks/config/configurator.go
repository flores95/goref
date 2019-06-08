package config

type Configurator interface {
	GetValue(string) string
	GetNamespace() string
	Load(map[string]string)
}

type Configurable interface {
	Configure(Configurator)
}

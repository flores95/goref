package config

type Configurator interface {
	GetValue(string) string
	Load(map[string]string)
}

type Configurable interface {
	Configure(Configurator)
}

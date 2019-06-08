package config

type BaseConfigurator struct {
	kvs       map[string]string
	namespace string
}

func NewBaseConfigurator(ns string) Configurator {
	c := new(BaseConfigurator)
	c.namespace = ns
	return c
}

func (c *BaseConfigurator) GetNamespace() string {
	return c.namespace
}

func (c *BaseConfigurator) GetValue(key string) (value string) {
	value = c.kvs[key]
	return value
}

func (c *BaseConfigurator) Load(kvs map[string]string) {
	c.kvs = kvs
}

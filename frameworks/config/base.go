package config

type BaseConfigurator struct {
	kvs       KVS
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
	if nsv := c.kvs[c.namespace+key]; nsv != "" {
		value = nsv
	}
	return value
}

func (c *BaseConfigurator) Load(kvs KVS) {
	c.kvs = kvs
}

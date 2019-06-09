package config

// BaseConfigurator is a simple key value store that can be used as the base for custom configurators
type BaseConfigurator struct {
	kvs       KVS
	namespace string
}

// NewBaseConfigurator create an instance of a base configurator
func NewBaseConfigurator(ns string) Configurator {
	c := new(BaseConfigurator)
	c.namespace = ns
	return c
}

// GetNamespace will return the namespace the configurator was initialied with
func (c *BaseConfigurator) GetNamespace() string {
	return c.namespace
}

// GetValue returns the value for the given key. If a namespace is set on the configurator that key will be used first.
func (c *BaseConfigurator) GetValue(key string) (value string) {
	value = c.kvs[key]
	if nsv := c.kvs[c.namespace+key]; nsv != "" {
		value = nsv
	}
	return value
}

// Load will load the configurator with a map (KVS) of keys and values
func (c *BaseConfigurator) Load(kvs KVS) {
	c.kvs = kvs
}

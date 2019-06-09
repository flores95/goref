package config

import (
	"github.com/joho/godotenv"
)

// DotenvConfigurator is a configurator that by default will retrieve its configurations from a .env file
type DotenvConfigurator struct {
	base Configurator
}

// NewDotenvConfigurator create a new configurator loading configuration data from a .env file
func NewDotenvConfigurator(ns string) Configurator {
	c := new(DotenvConfigurator)
	c.base = NewBaseConfigurator(ns)
	c.readEnv()
	return c
}

// GetNamespace returns the namespace this configurator was initialized with
func (c *DotenvConfigurator) GetNamespace() string {
	return c.base.GetNamespace()
}

// GetValue will return a value for the key provided if one exists
func (c *DotenvConfigurator) GetValue(key string) (value string) {
	value = c.base.GetValue(key)
	return value
}

// Load allows the manual loading/overiding of this configurators key value store
func (c *DotenvConfigurator) Load(kvs KVS) {
	c.base.Load(kvs)
}

func (c *DotenvConfigurator) readEnv() {
	kvs, _ := godotenv.Read() // will load setting from .env file
	c.base.Load(kvs)
}

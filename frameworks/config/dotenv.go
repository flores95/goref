package config

import (
	"github.com/joho/godotenv"
)

type DotenvConfigurator struct {
	base Configurator
}

func NewDotenvConfigurator(ns string) Configurator {
	c := new(DotenvConfigurator)
	c.base = NewBaseConfigurator(ns)
	c.readEnv()
	return c
}

func (c *DotenvConfigurator) GetNamespace() string {
	return c.base.GetNamespace()
}

func (c *DotenvConfigurator) GetValue(key string) (value string) {
	value = c.base.GetValue(key)
	return value
}

func (c *DotenvConfigurator) Load(kvs KVS) {
	c.base.Load(kvs)
}

func (c *DotenvConfigurator) readEnv() {
	kvs, _ := godotenv.Read() // will load setting from .env file
	c.base.Load(kvs)
}

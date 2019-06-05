package config

import (
	"github.com/joho/godotenv"
)

type DotenvConfigurator struct {
	kvs map[string]string
}

func NewDotenvConfigurator() Configurator {
	var c DotenvConfigurator = DotenvConfigurator{}
	c.kvs, _ = godotenv.Read() // will load setting from .env file
	return &c
}

func (c *DotenvConfigurator) GetValue(key string) (value string) {
	value = c.kvs[key]
	return value
}

func (c *DotenvConfigurator) Load(kvs map[string]string) {
	c.kvs = kvs
}

package config

import (
	"github.com/hashicorp/consul/api"
)

// KVClient is an interface for injecting a consul api KV client
type KVClient interface {
	KV() KVAccessor
}

// KVAccessor is an interface to allow mocking consul's api Client and KV struct
type KVAccessor interface {
	Get(key string, q *api.QueryOptions) (*api.KVPair, *api.QueryMeta, error)
	Put(p *api.KVPair, q *api.WriteOptions) (*api.WriteMeta, error)
}

// ConsulConfigurator is used to access a hashicorp consul key value store
type ConsulConfigurator struct {
	ns     string
	client *api.Client
}

// NewConsulConfigurator creates a new configurator connected to a consul service
func NewConsulConfigurator(ns string, client KVClient) Configurator {
	c := new(ConsulConfigurator)
	c.ns = ns
	if client == nil {
		c.consulInit()
	}
	return c
}

func (c *ConsulConfigurator) consulInit() {
	consulCfg := api.DefaultConfig()
	cc, _ := api.NewClient(consulCfg)
	c.client = cc
}

// GetNamespace returns the namespace this configurator was initialized with
func (c *ConsulConfigurator) GetNamespace() string {
	return c.ns
}

// GetValue will return a value for the key provided if one exists
func (c *ConsulConfigurator) GetValue(key string) string {
	kvp, _, _ := c.client.KV().Get(key, nil)
	return string(kvp.Value)
}

// Load allows the manual loading/overiding of this configurators key value store
func (c *ConsulConfigurator) Load(kvs KVS) {
	for k := range kvs {
		kvp := api.KVPair{Key: k, Value: []byte(kvs[k])}
		c.client.KV().Put(&kvp, nil)
	}
}

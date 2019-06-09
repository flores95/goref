package config

// MockConfigurator is a configurator that can be used for testing
type MockConfigurator struct {
	namespace       string
	kvs             KVS
	LoadInvoked     bool
	LoadFunc        func(KVS)
	GetValueInvoked bool
	GetValueFunc    func(string) string
}

// GetNamespace returns the namespace this configurator was initialized with
func (m *MockConfigurator) GetNamespace() string {
	return m.namespace
}

// GetValue will return a value for the key provided if one exists
func (m *MockConfigurator) GetValue(k string) string {
	m.GetValueInvoked = true
	return m.GetValueFunc(k)
}

// Load allows the manual loading/overiding of this configurators key value store
func (m *MockConfigurator) Load(kvs KVS) {
	m.LoadInvoked = true
	m.Load(kvs)
}

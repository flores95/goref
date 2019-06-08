package config

type MockConfigurator struct {
	namespace       string
	kvs             KVS
	LoadInvoked     bool
	LoadFunc        func(KVS)
	GetValueInvoked bool
	GetValueFunc    func(string) string
}

func (m *MockConfigurator) GetNamespace() string {
	return m.namespace
}

func (m *MockConfigurator) GetValue(k string) string {
	m.GetValueInvoked = true
	return m.GetValueFunc(k)
}

func (m *MockConfigurator) Load(kvs KVS) {
	m.LoadInvoked = true
	m.Load(kvs)
}

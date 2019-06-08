package config

type MockConfigurator struct {
	namespace       string
	kvs             map[string]string
	LoadInvoked     bool
	LoadFunc        func(map[string]string)
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

func (m *MockConfigurator) Load(kvs map[string]string) {
	m.LoadInvoked = true
	m.Load(kvs)
}

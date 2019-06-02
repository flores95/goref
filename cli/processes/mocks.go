package processes

type MockProcessor struct {
	DoInvoked bool
	DoFunc func()
	NameInvoked bool
	NameFunc func() string
}

func (m *MockProcessor) Do() {
	m.DoInvoked = true
	m.DoFunc()
}

func (m *MockProcessor) Name() string {
	m.NameInvoked = true
	return m.NameFunc()
}

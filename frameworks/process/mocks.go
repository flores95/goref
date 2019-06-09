package process

// MockProcessor is used for testing classes that require a processor
type MockProcessor struct {
	DoInvoked   bool
	DoFunc      func()
	NameInvoked bool
	NameFunc    func() string
}

// Do will set that it's been invoked and then executes the function provided to the mock
func (m *MockProcessor) Do() {
	m.DoInvoked = true
	m.DoFunc()
}

// Name will set that it's been invoked and then executes the function provided to the mock
func (m *MockProcessor) Name() string {
	m.NameInvoked = true
	return m.NameFunc()
}

package service

// MockServicer is used for testing classes that require a servicer
type MockServicer struct {
	RunInvoked    bool
	RunFunc       func()
	HealthInvoked bool
	HealthFunc    func() (bool, error)
	NameInvoked   bool
	NameFunc      func() string
}

// Run will set that it's been invoked and then execute the function provided to the mock
func (m *MockServicer) Run() {
	m.RunInvoked = true
	m.RunFunc()
}

// Health will set that it's been invoked and then execute the function provided to the mock
func (m *MockServicer) Health() (bool, error) {
	m.HealthInvoked = true
	return m.HealthFunc()
}

// Name will set that it's been invoked and then executes the function provided to the mock
func (m *MockServicer) Name() string {
	m.NameInvoked = true
	return m.NameFunc()
}

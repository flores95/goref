package log

import "github.com/flores95/goref/frameworks/config"

// MockLogger represents a logger for tests
type MockLogger struct {
	LogInvoked      bool
	LogFunc         func(Entry)
	SetLevelInvoked bool
	SetLevelFunc    func(Level)
}

// NewMockLogger creates a new mock logger for tests
func NewMockLogger(c config.Configurator) Logger {
	return new(MockLogger)
}

// Log flags that it was called and then executes the function provided to the mock
func (m *MockLogger) Log(e Entry) {
	m.LogInvoked = true
	m.LogFunc(e)
}

// SetLevel flags that it was called and then executes the function provided to the mock
func (m *MockLogger) SetLevel(l Level) {
	m.SetLevelInvoked = true
	m.SetLevelFunc(l)
}

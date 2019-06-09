package frameworks

import (
	"github.com/flores95/goref/frameworks/log"
)

// MockWorker defines a worker that can be used in test classes
type MockWorker struct {
	NameInvoked      bool
	NameFunc         func() string
	TypeInvoked      bool
	TypeFunc         func() string
	SetLoggerInvoked bool
	SetLoggerFunc    func(log.Logger)
	GetLoggerInvoked bool
	GetLoggerFunc    func() log.Logger
}

// Name returns the name of the worker
func (m *MockWorker) Name() string {
	m.NameInvoked = true
	return m.NameFunc()
}

// Type returns the string representation of the worker's type
func (m *MockWorker) Type() string {
	m.TypeInvoked = true
	return m.TypeFunc()
}

// SetLogger allows the worker's log to be set
func (m *MockWorker) SetLogger(l log.Logger) {
	m.SetLoggerInvoked = true
	m.SetLoggerFunc(l)
}

// GetLogger allows access to the workers logger
func (m *MockWorker) GetLogger() log.Logger {
	m.GetLoggerInvoked = true
	return m.GetLoggerFunc()
}

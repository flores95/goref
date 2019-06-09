package frameworks

import (
	"github.com/flores95/goref/frameworks/log"
)

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

func (m *MockWorker) Name() string {
	m.NameInvoked = true
	return m.NameFunc()
}

func (m *MockWorker) Type() string {
	m.TypeInvoked = true
	return m.TypeFunc()
}

func (m *MockWorker) SetLogger(l log.Logger) {
	m.SetLoggerInvoked = true
	m.SetLoggerFunc(l)
}

func (m *MockWorker) GetLogger() log.Logger {
	m.GetLoggerInvoked = true
	return m.GetLoggerFunc()
}

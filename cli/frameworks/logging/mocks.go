package logging

type MockLogger struct {
	LogInvoked      bool
	LogFunc         func(LogEvent)
	SetLevelInvoked bool
	SetLevelFunc    func(LogLevel)
}

func (m *MockLogger) Log(e LogEvent) {
	m.LogInvoked = true
	m.LogFunc(e)
}

func (m *MockLogger) SetLevel(l LogLevel) {
	m.SetLevelInvoked = true
	m.SetLevelFunc(l)
}

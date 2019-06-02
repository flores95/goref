package logging

type MockLogger struct {
	LogInvoked bool
	LogFunc    func(LogEvent)
}

func (m *MockLogger) Log(e LogEvent) {
	m.LogInvoked = true
	m.LogFunc(e)
}

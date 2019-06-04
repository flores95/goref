package logging

type FileLogger struct {
	fileName string
	level    LogLevel
}

func NewFileLogger(fn string) Logger {
	return FileLogger{fileName: fn}
}

func (l FileLogger) Log(le LogEvent) {
	if le.Level == Unassigned {
		le.Level = l.level
	}
}

func (l FileLogger) SetLevel(level LogLevel) {
	l.level = level
}

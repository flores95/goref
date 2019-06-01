package logging

type FileLogger struct {
	fileName string
}

func NewFileLogger(fn string) Logger {
	return FileLogger{fileName: fn}
}

func (l FileLogger) Log(le LogEvent) {
}

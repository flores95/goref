package logging

import "github.com/flores95/goref/frameworks/config"

type FileLogger struct {
	cfg      config.Configurator
	fileName string
	level    LogLevel
}

func NewFileLogger(c config.Configurator) Logger {
	l := new(FileLogger)
	l.cfg = c
	return l
}

func (l FileLogger) Log(le LogEvent) {
	if le.Level == Unassigned {
		le.Level = l.level
	}
}

func (l FileLogger) SetLevel(level LogLevel) {
	l.level = level
}

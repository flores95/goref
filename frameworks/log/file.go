package log

import "github.com/flores95/goref/frameworks/config"

// FileLogger represents a logger that writes logs to a file
type FileLogger struct {
	cfg      config.Configurator
	fileName string
	level    Level
}

// NewFileLogger creates a new file logger with the settings in the provided configurator
func NewFileLogger(c config.Configurator) Logger {
	l := new(FileLogger)
	l.cfg = c
	return l
}

// Log writes the log entry to the file
func (l FileLogger) Log(le Entry) {
	if le.Level == Unassigned {
		le.Level = l.level
	}
}

// SetLevel alow the level of logs written to be set
func (l FileLogger) SetLevel(level Level) {
	l.level = level
}

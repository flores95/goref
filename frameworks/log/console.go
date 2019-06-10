package log

import (
	"fmt"

	"github.com/flores95/goref/frameworks/config"
)

// ConsoleLogger will print logs to the console
type ConsoleLogger struct {
	cfg   config.Configurator
	level Level
}

// NewConsoleLogger creates a new console logger as configured by the given configurator
func NewConsoleLogger(c config.Configurator) Logger {
	l := new(ConsoleLogger)
	l.cfg = c
	return l
}

// Log will write the log entry to console
func (l *ConsoleLogger) Log(le Entry) {
	if le.Level == Unassigned {
		le.Level = l.level
	}
	fmt.Println(le)
}

// SetLevel allows the logging level to be set. This indicates which log entries to write
func (l *ConsoleLogger) SetLevel(level Level) {
	l.level = level
}

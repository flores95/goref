package log

import (
	"fmt"

	"github.com/flores95/goref/frameworks/config"
)

type ConsoleLogger struct {
	cfg   config.Configurator
	level LogLevel
}

func NewConsoleLogger(c config.Configurator) Logger {
	l := new(ConsoleLogger)
	l.cfg = c
	return l
}

func (l *ConsoleLogger) Log(le LogEvent) {
	if le.Level == Unassigned {
		le.Level = l.level
	}
	fmt.Println(le)
}

func (l *ConsoleLogger) SetLevel(level LogLevel) {
	l.level = level
}

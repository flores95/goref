package logging

import "fmt"

type ConsoleLogger struct {
	level LogLevel
}

func NewConsoleLogger() Logger {
	return ConsoleLogger{}
}

func (l ConsoleLogger) Log(le LogEvent) {
	if le.Level == Unassigned {
		le.Level = l.level
	}
	fmt.Println(le)
}

func (l ConsoleLogger) SetLevel(level LogLevel) {
	l.level = level
}

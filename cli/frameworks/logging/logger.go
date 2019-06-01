package logging

import (
	"fmt"
	"time"
)

type Logger interface {
	Log(LogEvent)
}

type LogEvent struct {
	Level   LogLevel
	Time    time.Time
	Context string
	Details string
}

type LogLevel int

const (
	Debug   LogLevel = 0
	Error   LogLevel = 1
	Warning LogLevel = 2
	Info    LogLevel = 3
)

func (level LogLevel) String() string {
	names := [...]string{
		"Debug",
		"Error",
		"Warning",
		"Info",
	}
	return names[level]
}

func (e LogEvent) String() string {
	return fmt.Sprintf("%v :: %v :: %v\n\t%v", e.Time, e.Level, e.Context, e.Details)
}

package logging

import (
	"fmt"
	"time"
)

type Logger interface {
	Log(LogEvent)
	SetLevel(LogLevel)
}

type Loggable interface {
	SetLogger(Logger)
	GetLogger() Logger
}

type LogEvent struct {
	Level   LogLevel
	Time    time.Time
	Context string
	Details string
}

type LogLevel int

const (
	Unassigned LogLevel = 0
	Debug      LogLevel = 1
	Error      LogLevel = 2
	Warning    LogLevel = 3
	Info       LogLevel = 4
)

func (level LogLevel) String() string {
	names := [...]string{
		"Unassigned",
		"Debug",
		"Error",
		"Warning",
		"Info",
	}
	return names[level]
}

func NewLogEvent(level LogLevel, ctx string, detail string) (event LogEvent) {
	event.Time = time.Now()
	event.Level = level
	event.Context = ctx
	event.Details = detail
	return event
}

func (e LogEvent) String() string {
	return fmt.Sprintf("%v :: %v :: %v\n\t%v", e.Time.Local(), e.Level, e.Context, e.Details)
}

package log

import (
	"fmt"
	"strings"
	"time"

	"github.com/flores95/goref/frameworks/config"
)

type (
	// Logger represents what an implementation of a logger requires
	Logger interface {
		Log(Entry)
		SetLevel(Level)
	}

	// Loggable represents what an implementation of a loggable class requires
	Loggable interface {
		SetLogger(Logger)
		GetLogger() Logger
	}

	// Entry represents a log entry
	Entry struct {
		Level   Level
		Time    time.Time
		Context string
		Details string
	}

	// Level is the level of logs to write
	Level int
)

// constants for log levels
const (
	Unassigned Level = 0
	Debug      Level = 1
	Error      Level = 2
	Warning    Level = 3
	Info       Level = 4
)

// NewLogger is a factory for the Logger interface
// It uses a configurator's LOG_USING key to decide what type of Logger to provide
func NewLogger(c config.Configurator) Logger {
	var l Logger
	lu := strings.ToUpper(c.GetValue("LOG_USING"))
	fmt.Println(lu)
	switch lu {
	case "CONSOLE":
		l = NewConsoleLogger(c)
	case "FILE":
		l = NewFileLogger(c)
	default:
		l = NewConsoleLogger(c)
	}
	return l
}

// NewLevel gets a Level from a string
func NewLevel(s string) Level {
	levels := map[string]Level{
		"unassigned": 0,
		"debug":      1,
		"error":      2,
		"warning":    3,
		"info":       4,
	}
	return levels[strings.ToLower(s)]
}

// String is the string representation of a log level
func (level Level) String() string {
	names := [...]string{
		"unassigned",
		"debug",
		"error",
		"warning",
		"info",
	}
	return names[level]
}

// NewEntry creates a new instance of a log event
func NewEntry(level Level, ctx string, detail string) (event Entry) {
	event.Time = time.Now()
	event.Level = level
	event.Context = ctx
	event.Details = detail
	return event
}

// String returns a string representation of a log event
func (e Entry) String() string {
	return fmt.Sprintf("%v :: %v :: %v\n\t%v", e.Time.Local(), e.Level, e.Context, e.Details)
}

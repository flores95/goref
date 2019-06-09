package log

import (
	"fmt"
	"strings"
	"time"

	"github.com/flores95/goref/frameworks/config"
)

type (
	Logger interface {
		Log(LogEvent)
		SetLevel(LogLevel)
	}

	Loggable interface {
		SetLogger(Logger)
		GetLogger() Logger
	}

	LogEvent struct {
		Level   LogLevel
		Time    time.Time
		Context string
		Details string
	}

	LogLevel int
)

const (
	Unassigned LogLevel = 0
	Debug      LogLevel = 1
	Error      LogLevel = 2
	Warning    LogLevel = 3
	Info       LogLevel = 4
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

func NewLogLevel(s string) LogLevel {
	levels := map[string]LogLevel{
		"unassigned": 0,
		"debug":      1,
		"error":      2,
		"warning":    3,
		"info":       4,
	}
	return levels[strings.ToLower(s)]
}

func (level LogLevel) String() string {
	names := [...]string{
		"unassigned",
		"debug",
		"error",
		"warning",
		"info",
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

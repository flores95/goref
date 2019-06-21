package frameworks

import (
	"github.com/flores95/goref/frameworks/log"
)

// Worker represents a framework tool ... like a logger or configurator or authenticator
type Worker interface {
	Namer
	Typer
	log.Loggable
}

// WorkerType defines a worker's type
type WorkerType int

// worker type constants
const (
	Config = 0
	Log    = 1
	Metric = 2
	Store  = 3
	Auth   = 4
)

func (wt WorkerType) String() string {
	names := [...]string{
		"Config",
		"Log",
		"Metric",
		"Store",
		"Auth",
	}
	return names[wt]
}

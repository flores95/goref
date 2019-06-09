package frameworks

import (
	"github.com/flores95/goref/frameworks/log"
)

type Worker interface {
	Nameable
	Typeable
	log.Loggable
}

type WorkerType int

const (
	Config  = 0
	Logging = 1
	Metrics = 2
	Storage = 3
	Auth    = 4
)

func (wt WorkerType) String() string {
	names := [...]string{
		"Config",
		"Logging",
		"Metrics",
		"Storage",
		"Auth",
	}
	return names[wt]
}

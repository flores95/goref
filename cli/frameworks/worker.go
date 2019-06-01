package frameworks

import (
	"github.com/flores95/golang-curriculum-c-5/cli/frameworks/logging"
)

type Worker interface {
	Nameable
	Typeable
	logging.Loggable
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

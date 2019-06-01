package frameworks

import "github.com/flores95/golang-curriculum-c-5/cli/frameworks/logging"

type Worker interface {
	GetName() string
	GetType() string
	logging.Loggable
}

type WorkerType int

const (
	Config  = 0
	Logging = 1
	Metrics = 2
	Storage = 3
)

func (wt WorkerType) String() string {
	names := [...]string{
		"Config",
		"Logging",
		"Metrics",
		"Storage",
	}
	return names[wt]
}

package processes

import "github.com/flores95/goref/cli/frameworks"

type Processor interface {
	Do()
	frameworks.Nameable
}

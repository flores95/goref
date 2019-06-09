package processes

import "github.com/flores95/goref/frameworks"

type Processor interface {
	Do()
	frameworks.Nameable
}

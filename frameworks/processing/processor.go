package processing

import "github.com/flores95/goref/frameworks"

type Processor interface {
	Do()
	frameworks.Nameable
}

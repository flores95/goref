package process

import "github.com/flores95/goref/frameworks"

// Processor represents a processor
type Processor interface {
	Do()
	frameworks.Nameable
}

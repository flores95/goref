package processes

import "github.com/flores95/golang-curriculum-c-5/cli/frameworks"

type Processor interface {
	Do()
	frameworks.Nameable
}

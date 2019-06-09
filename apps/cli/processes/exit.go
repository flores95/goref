package processes

import "github.com/flores95/goref/frameworks/process"

// ExitProcess manages the closing of the app
type ExitProcess struct {
	name string
}

// NewExitProcess creates a new exit processor
func NewExitProcess() process.Processor {
	proc := ExitProcess{}
	proc.name = "Exit"
	return proc
}

// Name returns the name of this processor
func (proc ExitProcess) Name() string {
	return proc.name
}

// Do executes the logic of this process
func (proc ExitProcess) Do() {
}

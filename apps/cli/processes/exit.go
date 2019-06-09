package processes

import "github.com/flores95/goref/frameworks/process"

type ExitProcess struct {
	name string
}

func NewExitProcess() process.Processor {
	proc := ExitProcess{}
	proc.name = "Exit"
	return proc
}

func (proc ExitProcess) Name() string {
	return proc.name
}

func (proc ExitProcess) Do() {
}

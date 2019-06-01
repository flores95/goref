package processes

type ExitProcess struct {
	name string
}

func NewExitProcess() Processor {
	proc := ExitProcess{}
	proc.name = "Exit"
	return proc
}

func (proc ExitProcess) GetName() string {
	return proc.name
}

func (proc ExitProcess) Do() {
}

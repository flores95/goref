package app

type ExitProcess struct {
	name string
	app App
}

func NewExitProcess(a App) Processor {
	proc := ExitProcess{}
	proc.name = "Exit"
	proc.app = a
	return proc
}

func (proc ExitProcess) GetName() string {
	return proc.name
}

func (proc ExitProcess) GetApp() App {
	return proc.app
}

func (proc ExitProcess) Do() {
	
}

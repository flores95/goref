package app

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/flores95/goref/frameworks/logging"
	"github.com/flores95/goref/cli/processes"
)

//App controls the application
type App struct {
	logger           logging.Logger
	interactiveProcs []processes.Processor
}

//NewApp creates a new application
func NewApp(
	l logging.Logger,
) (a App) {
	a.logger = l
	return a
}

func (a App) processCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}

	for _, p := range a.interactiveProcs {
		s = append(s, prompt.Suggest{Text: p.Name(), Description: ""})
	}

	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func (a App) GetProcessorByName(procs []processes.Processor, n string) (proc processes.Processor) {
	for _, p := range procs {
		if p.Name() == n {
			proc = p
			return proc
		}
	}
	return proc
}

func (a *App) RunInteractive(procs []processes.Processor) {
	a.interactiveProcs = procs

	for {
		fmt.Print(" What would you like to do now? ")
		p := a.GetProcessorByName(procs, prompt.Input("", a.processCompleter))
		p.Do()
		if p.Name() == "Exit" {
			break
		}
	}
}

func (a *App) RunAsync(procs []processes.Processor) {
	for _, p := range procs {
		go p.Do()
	}
}

func (a *App) RunInOrder(procs []processes.Processor) {
	for _, p := range procs {
		p.Do()
	}
}

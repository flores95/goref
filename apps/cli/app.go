package cli

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/flores95/goref/frameworks/log"
	"github.com/flores95/goref/frameworks/process"
)

//App controls the application
type App struct {
	logger           log.Logger
	interactiveProcs []process.Processor
}

//NewApp creates a new application
func NewApp(
	l log.Logger,
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

// GetProcessorByName finds the named processor in the given list
func (a App) GetProcessorByName(procs []process.Processor, n string) (proc process.Processor) {
	for _, p := range procs {
		if p.Name() == n {
			proc = p
			return proc
		}
	}
	return proc
}

// RunInteractive allows users to interactively select what to run from the given processors
func (a *App) RunInteractive(procs []process.Processor) {
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

// RunAsync will run all of the given processes asynchonously in separate go routines
func (a *App) RunAsync(procs []process.Processor) {
	for _, p := range procs {
		go p.Do()
	}
}

// RunInOrder runs the given processors in order synchonously
func (a *App) RunInOrder(procs []process.Processor) {
	for _, p := range procs {
		p.Do()
	}
}

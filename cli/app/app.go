package app

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/flores95/golang-curriculum-c-5/cli/frameworks"
	"github.com/flores95/golang-curriculum-c-5/cli/processes"
)

//App controls the application
type App struct {
	procs []processes.Processor
	fws   []frameworks.Worker
}

//NewApp creates a new application with injected controllers
func NewApp(
	procs []processes.Processor,
	fws []frameworks.Worker,
) (a App) {
	a.procs = procs
	a.fws = fws
	return a
}

func (a App) processCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}

	for _, p := range a.procs {
		s = append(s, prompt.Suggest{Text: p.GetName(), Description: ""})
	}

	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func (a App) GetProcessorByName(n string) (proc processes.Processor) {
	for _, p := range a.procs {
		if p.GetName() == n {
			proc = p
			return proc
		}
	}
	return proc
}

func (a *App) Run() {
	// fmt.Printf("Welcome, %v!\n", a.Users.GetCurrentUser())
	for {
		fmt.Print(" What would you like to do now? ")
		p := a.GetProcessorByName(prompt.Input("", a.processCompleter))
		p.Do()
		if p.GetName() == "Exit" {
			break
		}
	}
}

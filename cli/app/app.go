package app

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/flores95/golang-curriculum-c-5/cli/controllers"
)

//App controls the application
type App struct {
	Products controllers.ProductController
	Users    controllers.UserController
	Orders   controllers.OrderController
	procs    []Processor
}

//NewApp creates a new application with injected controllers
func NewApp(
	prod controllers.ProductController,
	user controllers.UserController,
	ord controllers.OrderController,
) (a App) {
	a.Products = prod
	a.Users = user
	a.Orders = ord
	return a
}

func (a App) processCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}

	for _, p := range a.procs {
		s = append(s, prompt.Suggest{Text: p.GetName(), Description: ""})
	}

	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func (a App) GetProcessorByName(n string) (proc Processor) {
	for _, p := range a.procs {
		if p.GetName() == n {
			proc = p
			return proc
		}
	}
	return proc
}

func (a *App) Run(procs []Processor) {
	a.procs = procs
	fmt.Printf("Welcome, %v!\n", a.Users.GetCurrentUser())
	for {
		fmt.Print(" What would you like to do now? ")
		p := a.GetProcessorByName(prompt.Input("", a.processCompleter))
		p.Do()
		if p.GetName() == "Exit" {
			break
		}
	}
}

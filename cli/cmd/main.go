package main

import (
	"github.com/flores95/goref/cli/app"
	"github.com/flores95/goref/cli/controllers"
	"github.com/flores95/goref/cli/processes"
	"github.com/flores95/goref/frameworks/logging"
)

func main() {
	// framework stuff
	l := logging.NewConsoleLogger()
	l.SetLevel(logging.Debug)

	// business logic handlers
	products := controllers.NewCLIProductController() // inject a framework data loader here (API / DB / JSON / CSV ???)
	users := controllers.NewCLIUserController()
	orders := controllers.NewCLIOrderController()

	//TODO move this into the constructor
	products.Load()
	users.Load()
	orders.Load()

	//TODO move this to a process
	var oProcs []processes.Processor
	oProcs = append(oProcs, processes.NewAuthenticateProcess(
		users,
	))

	// application flow processors
	var iProcs []processes.Processor
	iProcs = append(iProcs, processes.NewBuildOrderProcess(
		orders,
		products,
		users,
	))
	iProcs = append(iProcs, processes.NewExitProcess())

	// compose it all together and run
	a := app.NewApp(l)
	a.RunInOrder(oProcs)
	a.RunInteractive(iProcs)
}

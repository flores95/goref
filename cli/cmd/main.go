package main

import (
	"github.com/flores95/goref/cli/app"
	"github.com/flores95/goref/cli/controllers"
	"github.com/flores95/goref/cli/frameworks/logging"
	"github.com/flores95/goref/cli/processes"
)

func main() {
	// framework stuff
	logLevel := logging.Debug

	l := logging.NewConsoleLogger()
	l.Log(logging.NewLogEvent(logLevel, "MAIN", "Hey dude this works"))

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

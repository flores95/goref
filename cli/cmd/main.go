package main

import (
	"github.com/flores95/goref/cli/app"
	"github.com/flores95/goref/cli/controllers"
	"github.com/flores95/goref/cli/processes"
	"github.com/flores95/goref/frameworks/config"
	"github.com/flores95/goref/frameworks/logging"
)

func main() {
	// framework stuff
	c := config.NewDotenvConfigurator("CLI_")
	l := logging.NewLogger(c)
	ll := c.GetValue("LOG_LEVEL")
	l.SetLevel(logging.NewLogLevel(ll))

	// business logic handlers
	products := controllers.NewProductController(c)
	users := controllers.NewUserController(c)
	orders := controllers.NewOrderController(c)

	//TODO move this into the constructor
	products.Load()
	users.Load()
	orders.Load()

	// in order processes
	var oProcs []processes.Processor
	oProcs = append(oProcs, processes.NewAuthenticateProcess(
		*users,
	))

	// application flow processors
	var iProcs []processes.Processor
	iProcs = append(iProcs, processes.NewBuildOrderProcess(
		*orders,
		*products,
		*users,
	))
	iProcs = append(iProcs, processes.NewExitProcess())

	// compose it all together and run
	a := app.NewApp(l)
	a.RunInOrder(oProcs)
	a.RunInteractive(iProcs)
}

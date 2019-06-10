package main

import (
	"github.com/flores95/goref/apps/cli"
	"github.com/flores95/goref/apps/cli/processes"
	"github.com/flores95/goref/apps/controllers"
	"github.com/flores95/goref/frameworks/config"
	"github.com/flores95/goref/frameworks/log"
	"github.com/flores95/goref/frameworks/process"
)

func main() {
	// framework stuff
	c := config.NewDotenvConfigurator("CLI_")
	l := log.NewLogger(c)
	ll := c.GetValue("LOG_LEVEL")
	l.SetLevel(log.NewLevel(ll))

	// business logic handlers
	products := controllers.NewProductController(c)
	users := controllers.NewUserController(c)
	orders := controllers.NewOrderController(c)

	//TODO move this into the constructor
	products.Load()
	users.Load()
	orders.Load()

	// in order processes
	var oProcs []process.Processor
	oProcs = append(oProcs, processes.NewAuthenticateProcess(
		*users,
	))

	// application flow processors
	var iProcs []process.Processor
	iProcs = append(iProcs, processes.NewBuildOrderProcess(
		*orders,
		*products,
		*users,
	))
	iProcs = append(iProcs, processes.NewExitProcess())

	// compose it all together and run
	a := cli.NewApp(l)
	a.RunInOrder(oProcs)
	a.RunInteractive(iProcs)
}

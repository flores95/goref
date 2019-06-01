package main

import (
	"github.com/flores95/golang-curriculum-c-5/cli/app"
	"github.com/flores95/golang-curriculum-c-5/cli/controllers"
	"github.com/flores95/golang-curriculum-c-5/cli/frameworks"
	"github.com/flores95/golang-curriculum-c-5/cli/frameworks/logging"
	"github.com/flores95/golang-curriculum-c-5/cli/processes"
)

func main() {
	// framework stuff
	logLevel := logging.Debug

	l := logging.NewConsoleLogger()
	l.Log(logging.NewLogEvent(logLevel, "MAIN", "Hey dude this works"))
	fws := []frameworks.Worker{}

	// business logic handlers
	products := controllers.NewCLIProductController() // inject a framework data loader here (API / DB / JSON / CSV ???)
	users := controllers.NewCLIUserController()
	orders := controllers.NewCLIOrderController()

	//TODO move this into the constructor
	products.Load()
	users.Load()
	orders.Load()

	//TODO move this to a process
	users.SetCurrentUser()

	// application flow processors
	var procs []processes.Processor
	procs = append(procs, processes.NewBuildOrderProcess(
		orders,
		products,
		users,
	))
	procs = append(procs, processes.NewExitProcess())

	// compose it all together and run
	a := app.NewApp(procs, fws)
	a.Run()
}

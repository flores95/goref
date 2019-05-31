package main

import (
	"github.com/flores95/golang-curriculum-c-5/cli/app"
	"github.com/flores95/golang-curriculum-c-5/cli/controllers"
)

func main() {
	a := app.NewApp(
		controllers.NewCLIProductController(),
		controllers.NewCLIUserController(),
		controllers.NewCLIOrderController(),
	)
	a.Products.Load()
	a.Users.Load()
	a.Users.SetCurrentUser()

	var procs []app.Processor
	procs = append(procs, app.NewBuildOrderProcess(a))
	procs = append(procs, app.NewExitProcess(a))

	a.Run(procs)
}

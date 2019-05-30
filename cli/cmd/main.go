package main

import (
	"github.com/flores95/golang-curriculum-c-5/cli/app"
	"github.com/flores95/golang-curriculum-c-5/cli/controllers"
)

func main() {
	app := app.NewApp(
		controllers.NewCLIProductController(),
		controllers.NewCLIUserController(),
		controllers.NewCLIOrderController(),
	)
	app.Products.Load()
	app.Users.Load()
	app.Users.SetCurrentUser()
	app.BuildOrder()
}

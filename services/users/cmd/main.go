package main

import (
	"github.com/flores95/goref/frameworks/config"

	"github.com/flores95/goref/services/users"
)

func main() {
	conf := config.NewDotenvConfigurator("USERS_")

	svc := users.NewService(conf)
	svc.Start()

}

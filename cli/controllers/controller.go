package controllers

import (
	"github.com/flores95/goref/frameworks"
	"github.com/flores95/goref/frameworks/logging"
	"github.com/flores95/goref/cli/models"
)

type Controller interface {
	frameworks.Nameable
	logging.Loggable
}

//UserController interface defines the connection to a Product service
type UserController interface {
	Load()
	GetUserFromEmail(string) models.User
	GetAll() []models.User
	GetCurrentUser() models.User
	SetCurrentUser(models.User)
}

//ProductController interface defines the connection to a Product service
type ProductController interface {
	Load()
	GetAll() []models.Product
}

//OrderController interface defines the connection to a Product service
type OrderController interface {
	Load()
	PlaceOrder(models.Order) models.Order
}

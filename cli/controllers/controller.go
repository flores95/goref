package controllers

import (
	"github.com/flores95/golang-curriculum-c-5/cli/frameworks"
	"github.com/flores95/golang-curriculum-c-5/cli/frameworks/logging"
	"github.com/flores95/golang-curriculum-c-5/cli/models"
)

type Controller interface {
	frameworks.Nameable
	logging.Loggable
}

//UserController interface defines the connection to a Product service
type UserController interface {
	Load()
	GetCurrentUser() models.User
	SetCurrentUser()
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

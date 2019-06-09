package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/flores95/goref/apps/models"
	"github.com/flores95/goref/frameworks/config"
)

// UserController manages interactions and events for users
type UserController struct {
	users       []models.User
	currentUser models.User
	endpoint    string
	cfg         config.Configurator
}

// NewUserController creates a new user controller using defaults or provided configurator settings
func NewUserController(cfg config.Configurator) *UserController {
	c := new(UserController)
	c.cfg = cfg
	c.endpoint = "http://localhost:4180/users"
	if e := cfg.GetValue("USERS_EP"); e != "" {
		c.endpoint = e
	}
	return c
}

// Load retrieves all existing users
func (c *UserController) Load() {
	if len(c.users) > 0 {
		fmt.Println(":: Users loaded from cache")
		return
	}

	resp, err := http.Get("http://localhost:4180/users")
	if err != nil {
	}

	defer resp.Body.Close()

	var users []models.User
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("couldn't read from http")
	}
	json.Unmarshal(body, &users)

	c.users = users
}

// GetAll returns previously Load(ed) users
func (c UserController) GetAll() []models.User {
	return c.users
}

// GetUserFromEmail will return the first existing user to match the provide email
func (c UserController) GetUserFromEmail(e string) (user models.User) {
	for _, u := range c.users {
		if u.Email == e {
			user = u
			return user
		}
	}
	return user
}

// GetCurrentUser returns the currently authenticated user that is using the app
func (c *UserController) GetCurrentUser() models.User {
	return c.currentUser
}

// SetCurrentUser set the currently authenticated user that is using the app
func (c *UserController) SetCurrentUser(u models.User) {
	c.currentUser = u
}

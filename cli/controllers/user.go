package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/flores95/goref/cli/models"
)

type UserController struct {
	users       []models.User
	currentUser models.User
}

func NewUserController() *UserController {
	return &UserController{}
}

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

func (c UserController) GetAll() []models.User {
	return c.users
}

func (c UserController) GetUserFromEmail(e string) (user models.User) {
	for _, u := range c.users {
		if u.Email == e {
			user = u
			return user
		}
	}
	return user
}

func (c *UserController) GetCurrentUser() models.User {
	return c.currentUser
}

func (c *UserController) SetCurrentUser(u models.User) {
	c.currentUser = u
}

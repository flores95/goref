package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/c-bata/go-prompt"
	"github.com/flores95/golang-curriculum-c-5/cli/models"
)

type CLIUserController struct {
	users       []models.User
	currentUser models.User
}

func NewCLIUserController() *CLIUserController {
	return &CLIUserController{}
}

func (c *CLIUserController) Load() {
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

func (c CLIUserController) userFromEmail(e string) (user models.User) {
	for _, u := range c.users {
		if u.Email == e {
			user = u
			return user
		}
	}
	return user
}

func (c CLIUserController) completer(in prompt.Document) []prompt.Suggest {
	var s []prompt.Suggest

	for _, u := range c.users {
		s = append(s, prompt.Suggest{Text: u.Email, Description: u.Name})
	}

	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func (c *CLIUserController) GetCurrentUser() models.User {
	return c.currentUser
}

func (c *CLIUserController) SetCurrentUser() {
	c.currentUser = c.userFromEmail(prompt.Input("USER: ", c.completer))
}

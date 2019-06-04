package processes

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/flores95/goref/cli/controllers"
	"github.com/flores95/goref/cli/models"
)

type AuthenticateProcess struct {
	users controllers.UserController
	user  models.User
	name  string
}

func NewAuthenticateProcess(
	uc controllers.UserController,
) Processor {
	proc := AuthenticateProcess{}
	proc.name = "Authenicate User"
	proc.users = uc
	return proc
}

func (proc AuthenticateProcess) completer(in prompt.Document) []prompt.Suggest {
	var s []prompt.Suggest

	for _, u := range proc.users.GetAll() {
		s = append(s, prompt.Suggest{Text: u.Email, Description: u.Name})
	}

	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func (proc AuthenticateProcess) Name() string {
	return proc.name
}

func (proc AuthenticateProcess) Do() {
	var u models.User
	for {
		e := prompt.Input("USER: ", proc.completer)
		u = proc.users.GetUserFromEmail(e)
		if u.Email != "" {
			proc.user = u
			break
		}
	}

	fmt.Printf("Welcome %v!\n", proc.user.Name)
	proc.users.SetCurrentUser(proc.user)
}

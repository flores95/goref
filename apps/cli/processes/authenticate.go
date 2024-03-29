package processes

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/flores95/goref/apps/controllers"
	"github.com/flores95/goref/apps/models"
	"github.com/flores95/goref/frameworks/process"
)

// AuthenticateProcess authenticates a user and sets them as the current users
type AuthenticateProcess struct {
	users controllers.UserController
	user  models.User
	name  string
}

// NewAuthenticateProcess creates a authenticate processor 
func NewAuthenticateProcess(
	uc controllers.UserController,
) process.Processor {
	proc := AuthenticateProcess{}
	proc.name = "Authenicate User"
	proc.users = uc
	return proc
}

func (proc AuthenticateProcess) completer(in prompt.Document) []prompt.Suggest {
	var s []prompt.Suggest

	for _, u := range proc.users.GetAll() {
		s = append(s, prompt.Suggest{Text: u.Email(), Description: u.Name()})
	}

	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

// Name returns the name of the process
func (proc AuthenticateProcess) Name() string {
	return proc.name
}

// Do executes the processes logic 
func (proc AuthenticateProcess) Do() {
	var u models.User
	for {
		e := prompt.Input("USER: ", proc.completer)
		u = proc.users.GetUserFromEmail(e)
		if u.Email() != "" {
			proc.user = u
			break
		}
	}

	fmt.Printf("Welcome %v!\n", proc.user.Name())
	proc.users.SetCurrentUser(proc.user)
}

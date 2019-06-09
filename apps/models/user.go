package models

import (
	"fmt"
)

type User struct {
	Email string
	Name  string
	Phone string
}

func (u User) String() string {
	return fmt.Sprintf("%v :: %v :: %v", u.Name, u.Email, u.Phone)
}

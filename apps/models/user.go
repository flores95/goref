package models

import (
	"fmt"
)

// User represents a user of the application
type User struct {
	Email string
	Name  string
	Phone string
}

// String implements the Stringer interface and allows a product to be printed as a string
func (u User) String() string {
	return fmt.Sprintf("%v :: %v :: %v", u.Name, u.Email, u.Phone)
}

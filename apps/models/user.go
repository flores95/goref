package models

import (
	"encoding/json"
	"fmt"
)

// User represents a user of the application
type User struct {
	name  string
	email string
	phone string
}

// String implements the Stringer interface and allows a product to be printed as a string
func (u User) String() string {
	return fmt.Sprintf("%v :: %v :: %v", u.name, u.email, u.phone)
}

// MarshalJSON will create a JSON version of the user object including some of the invisible fields
func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"email": u.email,
		"name":  u.name,
		"phone": u.phone,
	})
}

// UnmarshalJSON allows json parsing to set invisible fields
func (u *User) UnmarshalJSON(b []byte) error {
	userJSON := struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}{}

	if err := json.Unmarshal(b, &userJSON); err != nil {
		return err
	}

	fmt.Println("AFTER UNMARSHAL")
	fmt.Println(userJSON.Email)

	u.email = userJSON.Email
	u.name = userJSON.Name
	u.phone = userJSON.Phone

	return nil
}

// Name provides access to a user's name
func (u *User) Name() string {
	return u.name
}

// Email provides access to a user's email address
func (u *User) Email() string {
	return u.email
}

// Phone provides access to a user's phone number
func (u *User) Phone() string {
	return u.phone
}

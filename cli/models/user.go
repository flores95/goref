package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Email string
	Name  string
	Phone string
}

func (u User) String() string {
	return fmt.Sprintf("%v :: %v :: %v", u.Name, u.Email, u.Phone)
}

func LoadUsers() []User {
	resp, err := http.Get("http://localhost:4180/users")
	if err != nil {
	}

	defer resp.Body.Close()

	var users []User
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("couldn't read from http")
	}
	json.Unmarshal(body, &users)

	return users
}

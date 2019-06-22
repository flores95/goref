package models

import (
	"github.com/flores95/goref/frameworks"
	"github.com/flores95/goref/frameworks/storage"
)

//User defines a user
type User struct {
	id    frameworks.ID
	tags  []frameworks.Tag
	email string
	name  string
	phone string
}

// NewUser is a constructor for a user model
func NewUser(id frameworks.ID, email string, name string, phone string) *User {
	u := new(User)
	u.id = id
	u.email = email
	u.name = name
	u.phone = phone
	return u
}

// ID provides access to the user's id as part of a storage item's interface
func (u *User) ID() frameworks.ID {
	return u.id
}

// Name provides access to the user's name as part of a storage item's interface
func (u *User) Name() string {
	return u.name
}

// Tags returns all the tags that have been added to this user
func (u *User) Tags() []frameworks.Tag {
	return u.tags
}

// HasTag indicates if this given tag has been added to this user
func (u *User) HasTag(t frameworks.Tag) bool {
	result := false
	return result
}

// LoadDemoUsers does what it says
func LoadDemoUsers(s storage.Store) {
	s.AddItem(NewUser(s.NextID(), "mike@fake.com", "Fake Mike", "NaN"))
	s.AddItem(NewUser(s.NextID(), "mike@test.com", "Test Mike", "123-456-7890"))
	s.AddItem(NewUser(s.NextID(), "user@test.com", "Test User", "555-555-5555"))
}

package repo

import "github.com/flores95/goref/services/users/user"

//Repo is the data store for the user microservice
type Repo struct {
	data []user.User
}

//Insert a user into the repository
func (r *Repo) Insert(u user.User) {
	r.data = append(r.data, u)
}

//GetAll users from the repository
func (r *Repo) GetAll() []user.User {
	return r.data
}

func (r *Repo) LoadDemoData() {
	r.Insert(user.User{Email: "mike@fake.com", Name: "Fake Mike", Phone: "NaN"})
	r.Insert(user.User{Email: "mike@test.com", Name: "Test Mike", Phone: "123-456-7890"})
	r.Insert(user.User{Email: "user@test.com", Name: "Test User", Phone: "555-555-5555"})
}

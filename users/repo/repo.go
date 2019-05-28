package repo

import "github.com/flores95/golang-curriculum-c-5/users/user"

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

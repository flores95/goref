package repo

import "github.com/flores95/goref/services/users/models"

//Repo is the data store for the user microservice
type Repo struct {
	data []models.User
}

//Insert a user into the repository
func (r *Repo) Insert(u models.User) {
	r.data = append(r.data, u)
}

//GetAll users from the repository
func (r *Repo) GetAll() []models.User {
	return r.data
}

// LoadDemoData does what it says
func (r *Repo) LoadDemoData() {
	r.Insert(models.User{Email: "mike@fake.com", Name: "Fake Mike", Phone: "NaN"})
	r.Insert(models.User{Email: "mike@test.com", Name: "Test Mike", Phone: "123-456-7890"})
	r.Insert(models.User{Email: "user@test.com", Name: "Test User", Phone: "555-555-5555"})
}

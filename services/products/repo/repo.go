package repo

import (
	"github.com/flores95/goref/services/products/models"
)

// Repo is the data store for the products microservice
type Repo struct {
	data   []models.Product
	nextID int
}

// Insert a product into the repository
func (r *Repo) Insert(p models.Product) models.Product {
	r.data = append(r.data, p)
	return p
}

// GetAll products from the repository
func (r *Repo) GetAll() []models.Product {
	return r.data
}

// LoadDemoData does what it says
func (r *Repo) LoadDemoData() {
	r.Insert(models.Product{UPC: "ITEM001", Name: "Nature's Miracle", Description: "Cleaning up after bad doggies"})
	r.Insert(models.Product{UPC: "ITEM002", Name: "Treats", Description: "Rewarding good doggies"})
	r.Insert(models.Product{UPC: "ITEM003", Name: "Tennis Ball", Description: "Wearing out energetic doggies"})
	r.Insert(models.Product{UPC: "ITEM004", Name: "Chow", Description: "Feeding hungry doggies"})
}

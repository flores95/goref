package repo

import (
	"github.com/flores95/goref/services/products/product"
)

//Repo is the data store for the products microservice
type Repo struct {
	data   []product.Product
	nextID int
}

//Insert a product into the repository
func (r *Repo) Insert(p product.Product) product.Product {
	r.data = append(r.data, p)
	return p
}

//GetAll products from the repository
func (r *Repo) GetAll() []product.Product {
	return r.data
}

func (r *Repo) LoadDemoData() {
	r.Insert(product.Product{UPC: "ITEM001", Name: "Nature's Miracle", Description: "Cleaning up after bad doggies"})
	r.Insert(product.Product{UPC: "ITEM002", Name: "Treats", Description: "Rewarding good doggies"})
	r.Insert(product.Product{UPC: "ITEM003", Name: "Tennis Ball", Description: "Wearing out energetic doggies"})
	r.Insert(product.Product{UPC: "ITEM004", Name: "Chow", Description: "Feeding hungry doggies"})
}

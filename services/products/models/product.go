package models

import (
	"github.com/flores95/goref/frameworks"
	"github.com/flores95/goref/frameworks/storage"
)

// Product represents products available for purchase
type Product struct {
	id    frameworks.ID
	tags  []frameworks.Tag
	name        string
	description string
}

// NewProduct is a constructor for a Product model
func NewProduct(id frameworks.ID, name string, description string) *Product {
	p := new(Product)
	p.id = id
	p.name = name
	p.description = description
	return p
}

// ID provides access to the product's id as part of a storage item's interface
func (p *Product) ID() frameworks.ID {
	return p.id
}

// Name provides access to the product's name as part of a storage item's interface
func (p *Product) Name() string {
	return p.name
}

// Description provides access to the product's descriptions
func (p *Product) Description() string {
	return p.name
}

// Tags returns all the tags that have been added to this product
func (p *Product) Tags() []frameworks.Tag {
	return p.tags
}

// HasTag indicates if this given tag has been added to this user
func (p *Product) HasTag(t frameworks.Tag) bool {
	result := false
	return result
}

// LoadDemoProducts does what it says
func LoadDemoProducts(s storage.Store) {
	s.AddItem(NewProduct(s.NextID(), "Nature's Miracle", "Cleaning up after bad doggies"))
	s.AddItem(NewProduct(s.NextID(), "Treats", "Rewarding good doggies"))
	s.AddItem(NewProduct(s.NextID(), "Tennis Ball", "Wearing out energetic doggies"))
	s.AddItem(NewProduct(s.NextID(), "Chow", "Feeding hungry doggies"))
}



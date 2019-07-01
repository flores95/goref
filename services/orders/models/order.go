package models

import (
	"fmt"

	"github.com/flores95/goref/frameworks"
)

type (
	// Order represents a stored customer order
	Order interface {
		frameworks.Identifier
		frameworks.Namer
		frameworks.Tagger
		Items() []Item
		UserEmail() string
		UserName() string
		UserPhone() string
	}

	// Item represents an item in an order
	Item interface {
		frameworks.Identifier
		Quantity() int
	}

	// order is a concrete implemntation of an Order
	order struct {
		id        frameworks.ID
		items     []item
		userEmail string
		userName  string
		userPhone string
		tags      []frameworks.Tag
	}

	// item is a concrete implementation of an item
	item struct {
		id       frameworks.ID
		quantity int
	}
)

// NewOrder will create a new order
func NewOrder() *Order {
	return new(Order)
}

// ID provides access to an order's ID
func (o *order) ID() frameworks.ID {
	return o.id
}

// Name provides access to an order name (email + id)
func (o *order) Name() string {
	return fmt.Sprintf("%v (%v)", o.id, o.userEmail)
}

// Items provides access to an order's items
func (o *order) Items() []item {
	return o.items
}

// UserEmail provides access to a user's email
func (o *order) UserEmail() string {
	return o.userEmail
}

// UserName provides access to a user's name
func (o *order) UserName() string {
	return o.userName
}

// UserPhone provides access to user's phonej
func (o *order) UserPhone() string {
	return o.userPhone
}

// Tags returns all the tags that have been added to this order
func (o *order) Tags() []frameworks.Tag {
	return o.tags
}

// HasTag indicates if this given tag has been added to this order
func (o *order) HasTag(t frameworks.Tag) bool {
	result := false
	return result
}

// Quantity provides the quantity of the item ordered
func (i *item) ID() frameworks.ID {
	return i.id
}

// Quantity provides the quantity of the item ordered
func (i *item) Quantity() int {
	return i.quantity
}

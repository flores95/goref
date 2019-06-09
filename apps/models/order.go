package models

import (
	"fmt"
)

// Order defines a customer's order
type Order struct {
	ID        string
	Items     []Item
	UserEmail string
	UserName  string
	UserPhone string
}

// Item is an item being ordered
type Item struct {
	ItemUPC  string
	Quantity int
}

// String implements the Stringer interface and allow an item to be printed as a string
func (i Item) String() string {
	return fmt.Sprintf("%d - %v", i.Quantity, i.ItemUPC)
}

// String implements the Stringer interface and allow an order to be printed as a string
func (o Order) String() (s string) {
	s += fmt.Sprintf("%v %v\n", o.ID, o.UserEmail)
	for _, i := range o.Items {
		s += fmt.Sprintln(i)
	}
	return s
}

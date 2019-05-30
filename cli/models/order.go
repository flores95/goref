package models

import (
	"fmt"
)

type Order struct {
	ID        string
	Items     []Item
	UserEmail string
	UserName  string
	UserPhone string
}

type Item struct {
	ItemUPC  string
	Quantity int
}

func (i Item) String() string {
	return fmt.Sprintf("%d - %v", i.Quantity, i.ItemUPC)
}

func (o Order) String() (s string) {
	s += fmt.Sprintf("%v %v\n", o.ID, o.UserEmail)
	for _, i := range o.Items {
		s += fmt.Sprintln(i)
	}
	return s
}

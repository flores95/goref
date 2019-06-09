package models

// Order represents a stored customer order
type Order struct {
	ID        string
	Items     []Item
	UserEmail string
	UserName  string
	UserPhone string
}

// Item represents an item in an order
type Item struct {
	ItemUPC  string
	Quantity int
}

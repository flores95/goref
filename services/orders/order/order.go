package order

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

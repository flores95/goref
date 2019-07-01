package models

// OrderSearch is just a struct to marshal the data for searching orders
type OrderSearch struct {
	userEmail string
}

// Find takes an array of orders and tries to find ones that match the search definition
func (os *OrderSearch) Find(orders []Order) (found []Order) {
	for _, o := range orders {
		if o.UserEmail() == os.userEmail {
			found = append(found, o)
		}
	}
	return found
}

package repo

import (
	"fmt"

	"github.com/flores95/golang-curriculum-c-5/orders/order"
)

//Repo is the data store for the orders microservice
type Repo struct {
	data   []order.Order
	nextID int
}

func (r *Repo) getNextID() string {
	if r.nextID == 0 {
		r.nextID = 1000
	}

	id := fmt.Sprintf("ORD-%d", r.nextID)
	r.nextID++
	return id
}

//Insert a orders into the repository
func (r *Repo) Insert(o order.Order) {
	o.ID = r.getNextID()
	r.data = append(r.data, o)
}

//GetAll orders from the repository
func (r *Repo) GetAll() []order.Order {
	return r.data
}

//GetOrders retrieves all orders for a given User Email
func (r *Repo) GetOrders(e string) (orders []order.Order) {
	for _, o := range r.data {
		if o.UserEmail == e {
			orders = append(orders, o)
		}
	}
	return orders
}

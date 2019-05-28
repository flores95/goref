package repo

import "github.com/flores95/golang-curriculum-c-5/orders/order"

//Repo is the data store for the orders microservice
type Repo struct {
	data []order.Order
}

//Insert a orders into the repository
func (r *Repo) Insert(u order.Order) {
	r.data = append(r.data, u)
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

package repo

import (
	"reflect"
	"testing"

	"github.com/flores95/golang-curriculum-c-5/orders/order"
)

func Test_Insert(t *testing.T) {
	r := Repo{}
	type args struct {
		o order.Order
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Can insert an order",
			args: args{order.Order{UserEmail: "insert@test.com"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.Insert(tt.args.o)
		})
	}
}

func TestRepo_GetAll(t *testing.T) {
	r := Repo{}
	r.Insert(order.Order{UserEmail: "user1@test.com"})
	r.Insert(order.Order{UserEmail: "user2@test.com"})
	r.Insert(order.Order{UserEmail: "user3@test.com"})
	tests := []struct {
		name string
		r    *Repo
		want []order.Order
	}{
		{
			name: "Get all orders",
			r:    &r,
			want: []order.Order{
				{UserEmail: "user1@test.com"},
				{UserEmail: "user2@test.com"},
				{UserEmail: "user3@test.com"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repo.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_GetOrders(t *testing.T) {
	r := Repo{}
	r.Insert(order.Order{UserEmail: "hasorders@test.com"})
	r.Insert(order.Order{UserEmail: "hasorders@test.com"})
	r.Insert(order.Order{UserEmail: "hasorders@test.com"})
	tests := []struct {
		name       string
		r          *Repo
		email      string
		wantOrders []order.Order
	}{
		{
			name:       "All orders for a user with orders",
			r:          &r,
			email:      "hasorders@test.com",
			wantOrders: r.GetAll(),
		},
		{
			name:       "All orders for a user with no orders",
			r:          &r,
			email:      "noorders@test.com",
			wantOrders: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOrders := tt.r.GetOrders(tt.email); !reflect.DeepEqual(gotOrders, tt.wantOrders) {
				t.Errorf("Repo.GetOrders() = %v, want %v", gotOrders, tt.wantOrders)
			}
		})
	}
}

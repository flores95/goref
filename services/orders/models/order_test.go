package models

import (
	"testing"

	"github.com/flores95/goref/frameworks"
)

func Test_OrderHasItems(t *testing.T) {
	o := order{
		items:     []item{item{id: "some-up-code", quantity: 42}},
		userEmail: "hasitems@test.com",
	}
	if len(o.items) != 1 {
		t.Error("Order has no items")
	}
}

func Test_OrderHasUserEmail(t *testing.T) {
	want := "test@testing.com"
	o := order{userEmail: "test@testing.com"}
	got := o.userEmail
	if got != want {
		t.Error("Order has no user email address")
	}
}

func Test_OrderHasId(t *testing.T) {
	want := "some id"
	o := order{id: "some id"}
	got := o.id
	if got != frameworks.ID(want) {
		t.Error("Order has no id")
	}
}

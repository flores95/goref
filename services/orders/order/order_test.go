package order

import (
	"testing"
)

func Test_OrderHasItems(t *testing.T) {
	o := Order{
		Items:     []Item{{ItemUPC: "some-up-code", Quantity: 42}},
		UserEmail: "hasitems@test.com",
	}
	if len(o.Items) != 1 {
		t.Error("Order has no items")
	}
}

func Test_OrderHasUserEmail(t *testing.T) {
	want := "test@testing.com"
	o := Order{UserEmail: "test@testing.com"}
	got := o.UserEmail
	if got != want {
		t.Error("Order has no user email address")
	}
}

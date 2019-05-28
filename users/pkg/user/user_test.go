package user

import (
	"testing"
)

func TestUserHasEmail(t *testing.T) {
	want := "test@testing.com"
	u := User{Email: "test@testing.com"}
	got := u.Email
	if got != want {
		t.Error("User has no Email")
	}
}

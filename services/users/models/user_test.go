package models

import (
	"testing"
)

func Test_UserHasEmail(t *testing.T) {
	want := "test@testing.com"
	u := User{email: "test@testing.com"}
	got := u.email
	if got != want {
		t.Error("User has no Email")
	}
}

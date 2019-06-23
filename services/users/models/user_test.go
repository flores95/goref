package models

import (
	"reflect"
	"testing"

	"github.com/flores95/goref/frameworks"
)

func Test_UserHasEmail(t *testing.T) {
	want := "test@testing.com"
	u := User{email: "test@testing.com"}
	got := u.email
	if got != want {
		t.Error("User has no Email")
	}
}

func TestNewUser(t *testing.T) {
	type args struct {
		id    frameworks.ID
		email string
		name  string
		phone string
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.id, tt.args.email, tt.args.name, tt.args.phone); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

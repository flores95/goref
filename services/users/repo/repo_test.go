package repo

import (
	"reflect"
	"testing"

	"github.com/flores95/golang-curriculum-c-5/users/user"
)

func Test_Insert(t *testing.T) {
	r := Repo{}
	type args struct {
		u user.User
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Can insert a user",
			args: args{user.User{Email: "insert@test.com"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.Insert(tt.args.u)
		})
	}
}

func TestRepo_GetAll(t *testing.T) {
	r := Repo{}
	r.Insert(user.User{Email: "user1@test.com"})
	r.Insert(user.User{Email: "user2@test.com"})
	r.Insert(user.User{Email: "user3@test.com"})
	tests := []struct {
		name string
		r    *Repo
		want []user.User
	}{
		{
			name: "Get all users",
			r:    &r,
			want: []user.User{
				{Email: "user1@test.com"},
				{Email: "user2@test.com"},
				{Email: "user3@test.com"},
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

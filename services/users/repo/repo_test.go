package repo

import (
	"reflect"
	"testing"

	"github.com/flores95/goref/services/users/models"
)

func Test_Insert(t *testing.T) {
	r := Repo{}
	type args struct {
		u models.User
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Can insert a user",
			args: args{models.User{Email: "insert@test.com"}},
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
	r.Insert(models.User{Email: "user1@test.com"})
	r.Insert(models.User{Email: "user2@test.com"})
	r.Insert(models.User{Email: "user3@test.com"})
	tests := []struct {
		name string
		r    *Repo
		want []models.User
	}{
		{
			name: "Get all users",
			r:    &r,
			want: []models.User{
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

package app

// import (
// 	"reflect"
// 	"testing"

// 	"github.com/flores95/golang-curriculum-c-5/cli/controllers"
// )

// func TestNewApp(t *testing.T) {
// 	type args struct {
// 		prod controllers.ProductController
// 		user controllers.UserController
// 		ord  controllers.OrderController
// 	}

// 	args = args{
// 		contollers.NewProduct, , Controller(),
// 	}
// 	tests := []struct {
// 		name  string
// 		args  args
// 		wantA App
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if gotA := NewApp(tt.args.prod, tt.args.user, tt.args.ord); !reflect.DeepEqual(gotA, tt.wantA) {
// 				t.Errorf("NewApp() = %v, want %v", gotA, tt.wantA)
// 			}
// 		})
// 	}
// }

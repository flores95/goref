package app

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/flores95/golang-curriculum-c-5/cli/controllers"
	"github.com/flores95/golang-curriculum-c-5/cli/models"
)

//App controls the application
type App struct {
	Products controllers.ProductController
	Users    controllers.UserController
	Orders   controllers.OrderController
}

//NewApp creates a new application with injected controllers
func NewApp(
	prod controllers.ProductController,
	user controllers.UserController,
	ord controllers.OrderController,
) (a App) {
	a.Products = prod
	a.Users = user
	a.Orders = ord
	return a
}

func emptyCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func (a App) productFromPrompt(prompt string) (product models.Product) {
	for _, p := range a.Products.GetAll() {
		if p.UPC == strings.Split(prompt, " :: ")[1] {
			product = p
			return product
		}
	}
	return product
}

func (a App) productsCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "x", Description: "Complete your order"},
	}

	for _, p := range a.Products.GetAll() {
		s = append(s, prompt.Suggest{Text: p.String(), Description: p.Description})
	}

	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

//BuildOrder walks a user through selecting products to order
func (a *App) BuildOrder() {
	u := a.Users.GetCurrentUser()
	fmt.Printf("Welcome %v, what would you like to order?\n", u.Name)

	var o models.Order
	for {
		sel := prompt.Input("ITEM: ", a.productsCompleter)
		if sel == "x" {
			break
		}
		item := a.productFromPrompt(sel)

		qs := prompt.Input(fmt.Sprintf("How many %v would you like?", item.Name), emptyCompleter)
		qty, _ := strconv.Atoi(qs)
		o.Items = append(o.Items, models.Item{ItemUPC: item.UPC, Quantity: qty})
	}
	fmt.Println("Your order:")
	fmt.Println(o)
	resp := prompt.Input("Would you like to place this order now?", emptyCompleter)

	if resp == "yes" {
		newOrder := a.Orders.PlaceOrder(o)
		fmt.Printf("Thank you %v, the following order has been placed for you:\n%v\n", u.Name, newOrder)
	}
}

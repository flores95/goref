package processes

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/flores95/goref/cli/controllers"
	"github.com/flores95/goref/cli/models"
)

type BuildOrderProcess struct {
	name     string
	orders   controllers.OrderController
	products controllers.ProductController
	users    controllers.UserController
}

func NewBuildOrderProcess(
	oc controllers.OrderController,
	pc controllers.ProductController,
	uc controllers.UserController,
) Processor {
	proc := BuildOrderProcess{}
	proc.name = "Order Products"
	proc.orders = oc
	proc.products = pc
	proc.users = uc
	return proc
}

func emptyCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func (proc BuildOrderProcess) Name() string {
	return proc.name
}

func (proc BuildOrderProcess) productFromPrompt(prompt string) (product models.Product) {
	for _, p := range proc.products.GetAll() {
		parsedInput := strings.Split(prompt, " :: ")
		if len(parsedInput) > 1 && p.UPC == parsedInput[1] {
			product = p
			return product
		}
	}
	return product
}

func (proc BuildOrderProcess) productsCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "x", Description: "Complete your order"},
	}

	for _, p := range proc.products.GetAll() {
		s = append(s, prompt.Suggest{Text: p.String(), Description: p.Description})
	}

	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func (proc BuildOrderProcess) Do() {
	u := proc.users.GetCurrentUser()
	fmt.Printf("%v, what would you like to order?\n", u.Name)

	var o models.Order
	for {
		sel := prompt.Input("ITEM: ", proc.productsCompleter)
		if sel == "x" {
			break
		}
		item := proc.productFromPrompt(sel)

		if item.UPC != "" {
			qs := prompt.Input(fmt.Sprintf("How many %v would you like?", item.Name), emptyCompleter)
			qty, _ := strconv.Atoi(qs)
			o.Items = append(o.Items, models.Item{ItemUPC: item.UPC, Quantity: qty})
		}
	}
	fmt.Println("Your order:")
	fmt.Println(o)
	resp := prompt.Input("Would you like to place this order now?", emptyCompleter)

	if resp == "yes" {
		newOrder := proc.orders.PlaceOrder(o)
		fmt.Printf("Thank you %v, the following order has been placed for you:\n%v\n", u.Name, newOrder)
	}
}

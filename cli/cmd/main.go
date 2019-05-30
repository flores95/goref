package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/flores95/golang-curriculum-c-5/cli/models"
)

var m struct {
	Products []models.Product
	Users    []models.User
	User     models.User
	Order    models.Order
}

func emptyCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func userCompleter(in prompt.Document) []prompt.Suggest {
	var s []prompt.Suggest

	for _, u := range m.Users {
		s = append(s, prompt.Suggest{Text: u.Email, Description: u.Name})
	}

	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func productCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "x", Description: "Complete your order"},
	}

	for _, p := range m.Products {
		s = append(s, prompt.Suggest{Text: p.String(), Description: p.Description})
	}

	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func userFromEmail(e string) (user models.User) {
	for _, u := range m.Users {
		if u.Email == e {
			user = u
			return user
		}
	}
	return user
}

func productFromPrompt(prompt string) (product models.Product) {
	for _, p := range m.Products {
		if p.UPC == strings.Split(prompt, " :: ")[1] {
			product = p
			return product
		}
	}
	return product
}

func main() {
	m.Products = models.LoadProducts()
	m.Users = models.LoadUsers()
	// m.User.Email = prompt.Input("EMAIL: ", emptyCompleter)
	// m.User.Name = prompt.Input("NAME: ", emptyCompleter)
	// m.User.Phone = prompt.Input("PHONE: ", emptyCompleter)
	m.User = userFromEmail(prompt.Input("USER: ", userCompleter))

	fmt.Printf("Welcome %v, what would you like to order?\n", m.User.Name)
	for {
		sel := prompt.Input("ITEM: ", productCompleter)
		if sel == "x" {
			break
		}
		item := productFromPrompt(sel)

		qs := prompt.Input(fmt.Sprintf("How many %v would you like?", item.Name), emptyCompleter)
		qty, _ := strconv.Atoi(qs)
		m.Order.Items = append(m.Order.Items, models.Item{ItemUPC: item.UPC, Quantity: qty})
	}
	fmt.Println("Your order:")
	fmt.Println(m.Order)
	resp := prompt.Input("Would you like to place this order now?", emptyCompleter)
	if resp == "yes" {
		m.Order.PlaceOrder()
		fmt.Println("Order placed.")
		fmt.Println(m.Order)
	}
}

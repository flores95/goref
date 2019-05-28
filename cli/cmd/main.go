package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/flores95/golang-curriculum-c-5/cli/models"
)

var m struct {
	Products []models.Product
	User models.User
	Order models.Order
}

func emptyCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
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

func main() {
	m.Products = models.LoadProducts()
	m.User.Email = prompt.Input("EMAIL: ", emptyCompleter)
	m.User.Name = prompt.Input("NAME: ", emptyCompleter)
	m.User.Phone = prompt.Input("PHONE: ", emptyCompleter)
	fmt.Printf("Welcome %v, what would you like to order?\n", m.User.Name)
	for {
		item := prompt.Input("ITEM: ", productCompleter)
		fmt.Println(item)
		if item == "x" {
			break
		}
	}
	fmt.Println(m.User)
}

package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Order struct {
	ID        string
	Items     []Item
	UserEmail string
	UserName  string
	UserPhone string
}

type Item struct {
	ItemUPC  string
	Quantity int
}

func (i Item) String() string {
	return fmt.Sprintf("%d - %v", i.Quantity, i.ItemUPC)
}

func (o Order) String() (s string) {
	s += fmt.Sprintf("%v %v\n", o.ID, o.UserEmail)
	for _, i := range o.Items {
		s += fmt.Sprintln(i)
	}
	return s
}

func (o *Order) PlaceOrder() {
	j, _ := json.Marshal(o)
	resp, _ := http.Post("http://localhost:4181/orders", "application/json", bytes.NewBuffer(j))

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result Order
	json.Unmarshal(body, &result)

	o.ID = result.ID
}

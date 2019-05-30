package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Product model for displaying available products
type Product struct {
	UPC         string
	Name        string
	Description string
}

func (p Product) String() string {
	return fmt.Sprintf("%v :: %v", p.Name, p.UPC)
}

func LoadProducts() []Product {
	resp, err := http.Get("http://localhost:4182/products")
	if err != nil {
	}

	defer resp.Body.Close()

	var products []Product
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("couldn't read from http")
	}
	json.Unmarshal(body, &products)

	return products
}

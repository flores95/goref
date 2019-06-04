package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/flores95/goref/cli/models"
)

type CLIProductController struct {
	products []models.Product
}

func NewCLIProductController() *CLIProductController {
	return &CLIProductController{}
}

func (c *CLIProductController) Load() {
	if len(c.products) > 0 {
		fmt.Println(":: Products loaded from cache")
		return
	}

	resp, err := http.Get("http://localhost:4182/products")
	if err != nil {
	}

	defer resp.Body.Close()

	var products []models.Product
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("couldn't read from http")
	}
	json.Unmarshal(body, &products)

	c.products = products
}

func (c *CLIProductController) GetAll() []models.Product {
	return c.products
}

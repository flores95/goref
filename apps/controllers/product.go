package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/flores95/goref/apps/models"
	"github.com/flores95/goref/frameworks/config"
)

type ProductController struct {
	products []models.Product
	endpoint string
	cfg      config.Configurator
}

func NewProductController(cfg config.Configurator) *ProductController {
	c := new(ProductController)
	c.cfg = cfg
	c.endpoint = "http://localhost:4181/products"
	if e := cfg.GetValue("PRODUCTS_EP"); e != "" {
		c.endpoint = e
	}
	return c
}

func (c *ProductController) Load() {
	if len(c.products) > 0 {
		fmt.Println(":: Products loaded from cache")
		return
	}

	resp, err := http.Get("http://localhost:4181/products")
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

func (c *ProductController) GetAll() []models.Product {
	return c.products
}

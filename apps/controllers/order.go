package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/flores95/goref/apps/models"
	"github.com/flores95/goref/frameworks/config"
)

// OrderController manages interactions with the order model
type OrderController struct {
	orders   []models.Order
	endpoint string
	cfg      config.Configurator
}

// NewOrderController creates a new order controller with eith default setting or setting passed throug the configurator
func NewOrderController(cfg config.Configurator) *OrderController {
	c := new(OrderController)
	c.cfg = cfg
	c.endpoint = "http://localhost:4182/orders"
	if e := cfg.GetValue("ORDERS_EP"); e != "" {
		c.endpoint = e
	}
	return c
}

// Load currently does nothing for orders but is necessary for the controller interface
func (c *OrderController) Load() {
	return
}

// PlaceOrder handles the placement of an order
func (c *OrderController) PlaceOrder(o models.Order) models.Order {
	var result models.Order

	j, _ := json.Marshal(o)

	if resp, err := http.Post("http://localhost:4182/orders", "application/json", bytes.NewBuffer(j)); err == nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &result)
	}

	return result
}

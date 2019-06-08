package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/flores95/goref/cli/models"
	"github.com/flores95/goref/frameworks/config"
)

type OrderController struct {
	orders   []models.Order
	endpoint string
	cfg      config.Configurator
}

func NewOrderController(cfg config.Configurator) *OrderController {
	c := new(OrderController)
	c.cfg = cfg
	c.endpoint = "http://localhost:4182/orders"
	if e := cfg.GetValue("ORDERS_EP"); e != "" {
		c.endpoint = e
	}
	return c
}

func (c *OrderController) Load() {
	return
}

func (c *OrderController) PlaceOrder(o models.Order) models.Order {
	j, _ := json.Marshal(o)
	resp, _ := http.Post("http://localhost:4182/orders", "application/json", bytes.NewBuffer(j))

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result models.Order
	json.Unmarshal(body, &result)

	return result
}

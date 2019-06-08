package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/flores95/goref/cli/models"
)

type OrderController struct {
	orders []models.Order
}

func NewOrderController() *OrderController {
	return &OrderController{}
}

func (c *OrderController) Load() {
	return
}

func (c *OrderController) PlaceOrder(o models.Order) models.Order {
	j, _ := json.Marshal(o)
	resp, _ := http.Post("http://localhost:4181/orders", "application/json", bytes.NewBuffer(j))

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result models.Order
	json.Unmarshal(body, &result)

	return result
}

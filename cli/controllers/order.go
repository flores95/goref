package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/flores95/goref/cli/models"
)

type CLIOrderController struct {
	orders []models.Order
}

func NewCLIOrderController() *CLIOrderController {
	return &CLIOrderController{}
}

func (c *CLIOrderController) Load() {
	return
}

func (c *CLIOrderController) PlaceOrder(o models.Order) models.Order {
	j, _ := json.Marshal(o)
	resp, _ := http.Post("http://localhost:4181/orders", "application/json", bytes.NewBuffer(j))

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result models.Order
	json.Unmarshal(body, &result)

	return result
}

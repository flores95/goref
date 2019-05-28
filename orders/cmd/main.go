package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/flores95/golang-curriculum-c-5/orders/order"
	"github.com/flores95/golang-curriculum-c-5/orders/repo"
)

var store = repo.Repo{}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/orders", allOrders).Methods("GET")
	r.HandleFunc("/orders", createOrder).Methods("POST")
	http.ListenAndServe(":4181", r)
}

func allOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(store.GetAll())
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u order.Order
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &u)
	store.Insert(u)
	j, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

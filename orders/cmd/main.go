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

type OrderSearch struct {
	UserEmail string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/orders", allOrders).Methods("GET")
	r.HandleFunc("/orders/find", findOrders).Methods("POST")
	r.HandleFunc("/orders", createOrder).Methods("POST")
	http.ListenAndServe(":4181", r)
}

func allOrders(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(store.GetAll())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func findOrders(w http.ResponseWriter, r *http.Request) {
	var os OrderSearch
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &os)
	j, _ := json.Marshal(store.GetOrders(os.UserEmail))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var o order.Order
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &o)
	newOrder := store.Insert(o)
	j, _ := json.Marshal(newOrder)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

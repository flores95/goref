package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/flores95/goref/frameworks/config"
	"github.com/flores95/goref/services/orders/order"
	"github.com/flores95/goref/services/orders/repo"
)

var store = repo.Repo{}

//OrderSearch is just a struct to marshal the data for searching orders
type OrderSearch struct {
	UserEmail string
}

func allOrders(w http.ResponseWriter, r *http.Request) {
	o := store.GetAll()
	j, _ := json.Marshal(o)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	fmt.Printf(":: [%v] ORDERS RETRIEVED ::\n", len(o))
}

func findOrders(w http.ResponseWriter, r *http.Request) {
	var os OrderSearch
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &os)
	o := store.GetOrders(os.UserEmail)
	j, _ := json.Marshal(o)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	fmt.Printf(":: [%v] ORDERS FOUND ::\n", len(o))
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
	fmt.Printf(":: ORDER CREATED :: [%v]\n", newOrder.ID)
}

func main() {
	l := config.NewDotenvConfigurator("ORDERS_")
	fmt.Printf(":: ORDERS MICROSERVICE :: http://localhost:%v\n", l.GetValue("PORT"))
	r := mux.NewRouter()
	r.HandleFunc("/orders", allOrders).Methods("GET")
	r.HandleFunc("/orders/find", findOrders).Methods("POST")
	r.HandleFunc("/orders", createOrder).Methods("POST")
	http.ListenAndServe(":"+l.GetValue("PORT"), r)
}

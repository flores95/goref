package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/flores95/goref/frameworks/config"
	"github.com/flores95/goref/frameworks/storage"
	"github.com/flores95/goref/services/orders/models"
)

var store = storage.NewMemoryStore("orders", "OID", 1000)

func allOrders(w http.ResponseWriter, r *http.Request) {
	o := store.Items()
	j, _ := json.Marshal(o)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	fmt.Printf(":: [%v] ORDERS RETRIEVED ::\n", len(o))
}

func findOrders(w http.ResponseWriter, r *http.Request) {
	var os models.OrderSearch
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &os)

	var orders []models.Order
	for _, i := range store.Items() {
		orders = append(orders, i.(models.Order))
	}

	o := os.Find(orders)
	j, _ := json.Marshal(o)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	fmt.Printf(":: [%v] ORDERS FOUND ::\n", len(o))
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var o models.Order
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &o)
	store.AddItem(o)
	j, _ := json.Marshal(&o)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	fmt.Printf(":: ORDER CREATED :: [%v]\n", o.ID())
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

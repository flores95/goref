package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/flores95/goref/frameworks/config"
	"github.com/flores95/goref/frameworks/storage"
	"github.com/flores95/goref/services/products/models"
)

var store = storage.NewMemoryStore("products", "UPC", 10000)

func allProducts(w http.ResponseWriter, r *http.Request) {
	p := store.Items()
	j, _ := json.Marshal(p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	fmt.Printf(":: [%v] PRODUCTS RETRIEVED ::\n", len(p))
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &p)
	store.AddItem(&p)
	j, _ := json.Marshal(&p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	fmt.Printf(":: PRODUCT ADDED :: [%v]\n", p)
}

func main() {
	l := config.NewDotenvConfigurator("PRODUCTS_")
	fmt.Printf(":: PRODUCTS MICROSERVICE :: http://localhost:%v\n", l.GetValue("PORT"))
	models.LoadDemoProducts(store)
	fmt.Printf(":: [%v] DEMO PRODUCTS LOADED ::\n", len(store.Items()))
	r := mux.NewRouter()
	r.HandleFunc("/products", allProducts).Methods("GET")
	r.HandleFunc("/products", createProduct).Methods("POST")
	http.ListenAndServe(":"+l.GetValue("PORT"), r)
}

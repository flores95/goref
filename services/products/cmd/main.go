package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/flores95/goref/frameworks/config"
	"github.com/flores95/goref/services/products/models"
	"github.com/flores95/goref/services/products/repo"
)

var store = repo.Repo{}

func allProducts(w http.ResponseWriter, r *http.Request) {
	p := store.GetAll()
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
	newProduct := store.Insert(p)
	j, _ := json.Marshal(newProduct)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	fmt.Printf(":: PRODUCT ADDED :: [%v]\n", p)
}

func main() {
	l := config.NewDotenvConfigurator("PRODUCTS_")
	fmt.Printf(":: PRODUCTS MICROSERVICE :: http://localhost:%v\n", l.GetValue("PORT"))
	store.LoadDemoData()
	fmt.Printf(":: [%v] DEMO PRODUCTS LOADED ::\n", len(store.GetAll()))
	r := mux.NewRouter()
	r.HandleFunc("/products", allProducts).Methods("GET")
	r.HandleFunc("/products", createProduct).Methods("POST")
	http.ListenAndServe(":"+l.GetValue("PORT"), r)
}

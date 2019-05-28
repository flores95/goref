package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/flores95/golang-curriculum-c-5/products/product"
	"github.com/flores95/golang-curriculum-c-5/products/repo"
)

var store = repo.Repo{}

func allProducts(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(store.GetAll())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var p product.Product
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &p)
	newProduct := store.Insert(p)
	j, _ := json.Marshal(newProduct)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func main() {
	store.LoadDemoData()
	r := mux.NewRouter()
	r.HandleFunc("/products", allProducts).Methods("GET")
	r.HandleFunc("/products", createProduct).Methods("POST")
	http.ListenAndServe(":4182", r)
}

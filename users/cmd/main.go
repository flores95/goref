package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/flores95/golang-curriculum-c-5/users/repo"
	"github.com/flores95/golang-curriculum-c-5/users/user"
)

var store = repo.Repo{}

func allUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u := store.GetAll()
	j, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	fmt.Printf(":: [%v] USERS RETRIEVED ::\n", len(u))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u user.User
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &u)
	store.Insert(u)
	j, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	fmt.Printf(":: USER CREATED :: [%v]\n", u)
}

func main() {
	fmt.Println(":: USERS MICROSERVICE :: http://localhost:4180")
	store.LoadDemoData()
	fmt.Printf(":: [%v] DEMO USERS LOADED ::\n", len(store.GetAll()))
	r := mux.NewRouter()
	r.HandleFunc("/users", allUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	http.ListenAndServe(":4180", r)
}

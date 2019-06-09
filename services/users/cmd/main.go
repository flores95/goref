package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/flores95/goref/frameworks/config"

	"github.com/gorilla/mux"

	"github.com/flores95/goref/services/users/models"
	"github.com/flores95/goref/services/users/repo"
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
	var u models.User
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &u)
	store.Insert(u)
	j, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	fmt.Printf(":: USER CREATED :: [%v]\n", u)
}

func main() {
	l := config.NewDotenvConfigurator("USERS_")
	fmt.Printf(":: USERS MICROSERVICE :: http://localhost:%v\n", l.GetValue("PORT"))
	store.LoadDemoData()
	fmt.Printf(":: [%v] DEMO USERS LOADED ::\n", len(store.GetAll()))
	r := mux.NewRouter()
	r.HandleFunc("/users", allUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	http.ListenAndServe(":"+l.GetValue("PORT"), r)
}

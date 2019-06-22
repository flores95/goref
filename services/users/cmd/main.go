package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/flores95/goref/frameworks/config"
	"github.com/flores95/goref/frameworks/storage"

	"github.com/gorilla/mux"

	"github.com/flores95/goref/services/users/models"
)

var store = storage.NewMemoryStore("users", "UID", 1000)

func allUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u := store.Items()
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
	store.AddItem(&u)
	j, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	fmt.Printf(":: USER CREATED :: [%v]\n", u)
}

func main() {
	l := config.NewDotenvConfigurator("USERS_")
	fmt.Printf(":: USERS MICROSERVICE :: http://localhost:%v\n", l.GetValue("PORT"))
	models.LoadDemoUsers(store)
	fmt.Printf(":: [%v] DEMO USERS LOADED ::\n", len(store.Items()))
	r := mux.NewRouter()
	r.HandleFunc("/users", allUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	http.ListenAndServe(":"+l.GetValue("PORT"), r)
}

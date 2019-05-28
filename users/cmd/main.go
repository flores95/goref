package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/flores95/golang-curriculum-c-5/users/repo"
	"github.com/flores95/golang-curriculum-c-5/users/user"
)

var store = repo.Repo{}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", allUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	http.ListenAndServe(":4180", r)
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(store.GetAll())
	w.WriteHeader(http.StatusOK)
	w.Write(j)
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
}

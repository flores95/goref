package users

import (
	"fmt"
	"net/http"

	"github.com/flores95/goref/frameworks/config"
	"github.com/flores95/goref/frameworks/log"
	"github.com/flores95/goref/frameworks/service"
	"github.com/flores95/goref/frameworks/storage"
	"github.com/flores95/goref/services/users/models"
)

// Service provides access to and management of one or more servicers or endpoints
type Service struct {
	base    service.HTTPServicer
	config  config.Configurator
	log     log.Logger
	store   storage.Store
	handler Handler
}

// NewService uses the injected configurator to create a new Service
func NewService(c config.Configurator) (svc Service) {
	svc.config = c
	svc.log = log.NewLogger(c)
	svc.store = storage.NewStore(c)

	svc.base = service.NewHTTPServicer(svc.config)
	svc.handler = NewHandler(svc.store)

	return svc
}

func (svc *Service) allUsers(w http.ResponseWriter, r *http.Request) {
	result := svc.handler.All()
	svc.base.Respond(w, result, http.StatusOK)
}

func (svc *Service) createUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	svc.base.Read(r, &u)

	result := svc.handler.Create(u)

	svc.base.Respond(w, result, http.StatusOK)
}

// Start the service
func (svc *Service) Start() {
	svc.log.Log(log.NewEntry(
		log.NewLevel("info"),
		"USERS MICROSERVICE",
		fmt.Sprintf(":: STARTING SERVICES ON %v", svc.config.GetValue("PORT")),
	))

	models.LoadDemoUsers(svc.store)
	svc.log.Log(log.NewEntry(
		log.NewLevel("info"),
		"USERS MICROSERVICE",
		fmt.Sprintf(":: [%v] DEMO USERS LOADED ::\n", len(svc.store.Items())),
	))

	svc.base.Router.HandleFunc("/users", svc.allUsers).Methods("GET")
	svc.base.Router.HandleFunc("/users", svc.createUser).Methods("POST")

	svc.base.Start()
}

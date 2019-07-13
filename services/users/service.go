package users

import (
	"github.com/flores95/goref/frameworks/config"
	"github.com/flores95/goref/frameworks/log"
	"github.com/flores95/goref/frameworks/storage"
)

// Service provides access to and management of one or more servicers or endpoints
type Service struct{
	config config.Configurator
	log log.Logger
	store storage.Store
}

// NewService uses the injected configurator to create a new Service
func NewService(c config.Configurator) (svc Service) {
	svc.config = c
	svc.log = log.NewLogger(c)
	svc.store = storage.NewStore(c)
	return svc
}
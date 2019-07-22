package users

import (
	"github.com/flores95/goref/frameworks/service"
	"github.com/flores95/goref/frameworks/storage"

	"github.com/flores95/goref/services/users/models"
)

// Handler is used by services to execute business logic
type Handler struct {
	store storage.Store
}

// NewHandler will create a new handler with the given configuration
func NewHandler(s storage.Store) (handler Handler) {
	handler.store = s
	return handler
}

// All will retrieve all users
func (h Handler) All() (results []models.User) {
	for _, i := range h.store.Items() {
		results = append(results, i.(models.User))
	}
	return results
}

// Create will create a new user
func (h Handler) Create(user models.User) models.User {
	h.store.AddItem(&user)
	return user
}

// Health can be called to verify this handler has access to what it needs
func (h Handler) Health() (s service.Status, e error) {
	return service.Up, nil
}

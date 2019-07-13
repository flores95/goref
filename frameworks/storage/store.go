package storage

import (
	"strconv"
	"strings"

	"github.com/flores95/goref/frameworks"
	"github.com/flores95/goref/frameworks/config"
)

type (
	// Store represents a data store
	Store interface {
		frameworks.Namer
		NextID() frameworks.ID
		AddItem(Item)
		Items() []Item
		GetItem(frameworks.ID) Item
	}

	// Item represents an item that can be stored
	Item interface {
		frameworks.Identifier
		frameworks.Namer
		frameworks.Tagger
	}
)

// NewStore is a factory for the Store interface
// It uses a configurator's STORE_USING key to decide what type of store to provide
func NewStore(c config.Configurator) (s Store) {
	using := strings.ToUpper(c.GetValue("STORE_USING"))
	name := c.GetValue("STORE_NAME")
	idPrefix := c.GetValue("STORE_ID_PREFIX")
	idSeed, _ := strconv.Atoi(c.GetValue("STORE_ID_SEED"))

	switch using {
	case "MEMORY":
		s = NewMemoryStore(name, idPrefix, idSeed)
	default:
		s = NewMemoryStore(name, idPrefix, idSeed)
	}
	return s
}

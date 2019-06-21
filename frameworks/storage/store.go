package storage

import (
	"github.com/flores95/goref/frameworks"
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

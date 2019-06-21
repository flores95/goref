package storage

import (
	"fmt"

	"github.com/flores95/goref/frameworks"
)

type (
	// MemoryStore is memory based storage
	MemoryStore struct {
		name     string
		items    []Item
		nextID   int
		idPrefix string
	}
)

// NewMemoryStore is a MemoryStore constructor
func NewMemoryStore(name string, idPrefix string, idSeed int) *MemoryStore {
	s := new(MemoryStore)
	s.name = name
	s.idPrefix = idPrefix
	s.nextID = idSeed
	return s
}

// Name returns the name of the store and is part of the Namer interface
func (s *MemoryStore) Name() string {
	return s.name
}

// NextID returns the next available id or a store item
func (s *MemoryStore) NextID() frameworks.ID {
	id := fmt.Sprintf("%v-%d", s.idPrefix, s.nextID)
	s.nextID++
	return frameworks.ID(id)
}

// AddItem adds an item to the memory store
func (s *MemoryStore) AddItem(i Item) {
	s.items = append(s.items, i)
}

// Items returns all the items in the store
func (s *MemoryStore) Items() []Item {
	return s.items
}

// GetItem returns the item matching the provided id
func (s *MemoryStore) GetItem(id frameworks.ID) (i Item) {
	for _, item := range s.items {
		if item.ID() == id {
			return item
		}
	}
	return i
}

package storage

import (
	"reflect"
	"testing"

	"github.com/flores95/goref/frameworks"
)

func TestNewMemoryStore(t *testing.T) {
	tests := []struct {
		name      string
		storeName string
		idPrefix  string
		idSeed    int
		wantID    frameworks.ID
	}{
		{
			name:      "Create a memory store with a simple name, prefix and id seed",
			storeName: "test-store",
			idPrefix:  "TESTID",
			idSeed:    1000,
			wantID:    "TESTID-1000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewMemoryStore(tt.storeName, tt.idPrefix, tt.idSeed)
			gotID := s.NextID()
			if !reflect.DeepEqual(gotID, tt.wantID) {
				t.Errorf("%v :: got = %v :: want %v", tt.name, gotID, tt.wantID)
			}
			gotName := s.Name()
			if !reflect.DeepEqual(gotName, tt.storeName) {
				t.Errorf("%v :: got = %v :: want %v", tt.name, gotName, tt.storeName)
			}
		})
	}
}

func TestMemoryStore_AddItem(t *testing.T) {
	mi := new(MockItem)
	mi.IDFunc = func() frameworks.ID { return "testid-1000" }
	tests := []struct {
		name string
		s    *MemoryStore
		i    Item
	}{
		{name: "Add Item to Memory Store", s: NewMemoryStore("test-store", "testid", 1000), i: mi},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.AddItem(tt.i)
			got := tt.s.GetItem(tt.i.ID())
			if !reflect.DeepEqual(got, tt.i) {
				t.Errorf("%v :: got = %v :: want %v", tt.name, got, tt.i)
			}
		})
	}
}

func TestMemoryStore_Items(t *testing.T) {
	ms := NewMemoryStore("test-stor", "testid", 1000)
	var mockItems []Item
	for i := 0; i < 5; i++ {
		mi := new(MockItem)
		id := ms.NextID()
		mi.IDFunc = func() frameworks.ID { return id }
		mockItems = append(mockItems, mi)
		ms.AddItem(mi)
	}
	tests := []struct {
		name string
		s    *MemoryStore
		want []Item
	}{
		{name: "Can get a list of all the items in a store", s: ms, want: mockItems},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Items(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v :: got = %v :: want %v", tt.name, got, tt.want)
			}
		})
	}
}

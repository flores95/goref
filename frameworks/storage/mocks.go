package storage

import (
	"github.com/flores95/goref/frameworks"
)

type (
	// MockItem represents a storage Item for testing
	MockItem struct {
		IDInvoked     bool
		IDFunc        func() frameworks.ID
		NameInvoked   bool
		NameFunc      func() string
		TagsInvoked   bool
		TagsFunc      func() []frameworks.Tag
		HasTagInvoked bool
		HasTagFunc    func(frameworks.Tag) bool
	}
)

// ID is mocked function that is a pass through for MockItem.ID()
func (i *MockItem) ID() frameworks.ID {
	i.IDInvoked = true
	return i.IDFunc()
}

// Name is mocked function that is a pass through for MockItem.Name()
func (i *MockItem) Name() string {
	i.NameInvoked = true
	return i.NameFunc()
}

// Tags is mocked function that is a pass through for MockItem.Tags()
func (i *MockItem) Tags() []frameworks.Tag {
	i.TagsInvoked = true
	return i.TagsFunc()
}

// HasTag is mocked function that is a pass through for MockItem.HasTag()
func (i *MockItem) HasTag(t frameworks.Tag) bool {
	i.HasTagInvoked = true
	return i.HasTagFunc(t)
}

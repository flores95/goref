package models

import (
	"fmt"
)

// Product represents a product that can be sold/ordered
type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description`
}

// String implements the Stringer interface and allows a product to be printed as a string
func (p Product) String() string {
	return fmt.Sprintf("%v :: %v", p.Name, p.ID)
}

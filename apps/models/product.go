package models

import (
	"fmt"
)

// Product represents a product that can be sold/ordered
type Product struct {
	UPC         string
	Name        string
	Description string
}

// String implements the Stringer interface and allows a product to be printed as a string
func (p Product) String() string {
	return fmt.Sprintf("%v :: %v", p.Name, p.UPC)
}

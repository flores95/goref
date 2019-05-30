package models

import (
	"fmt"
)

//Product model for displaying available products
type Product struct {
	UPC         string
	Name        string
	Description string
}

func (p Product) String() string {
	return fmt.Sprintf("%v :: %v", p.Name, p.UPC)
}

package models

import (
	"fmt"
)

type Product struct {
	UPC         string
	Name        string
	Description string
}

func (p Product) String() string {
	return fmt.Sprintf("%v :: %v", p.Name, p.UPC)
}

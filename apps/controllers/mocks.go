package controllers

import "github.com/flores95/goref/apps/models"

// MockProductController creates a controller to use for tests
type MockProductController struct {
}

// GetAll does nothing for the mock controller
func (m *MockProductController) GetAll() []models.Product {
	return nil
}

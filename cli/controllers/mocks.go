package controllers

import "github.com/flores95/goref/cli/models"

type MockProductController struct {
}

func (m *MockProductController) GetAll() []models.Product {
	return nil
}

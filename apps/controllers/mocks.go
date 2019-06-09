package controllers

import "github.com/flores95/goref/apps/models"

type MockProductController struct {
}

func (m *MockProductController) GetAll() []models.Product {
	return nil
}

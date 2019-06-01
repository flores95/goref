package controllers

import "github.com/flores95/golang-curriculum-c-5/cli/models"

type MockProductController struct {
}

func (m *MockProductController) GetAll() []models.Product {
	return nil
}

package repository

import (
	"github.com/manuelfirman/go-API/internal"
	"github.com/stretchr/testify/mock"
)

// NewProductMock creates a new ProductMock
func NewProductMock() *ProductMock {
	return &ProductMock{}
}

// ProductMock is a mock type for the ProductRepository interface
type ProductMock struct {
	mock.Mock
}

// GetAll mocks the GetAll method
func (m *ProductMock) GetAll() ([]internal.Product, error) {
	args := m.Called()
	return args.Get(0).([]internal.Product), args.Error(1)
}

// GetByID mocks the GetByID method
func (m *ProductMock) Get(id int) (internal.Product, error) {
	args := m.Called(id)
	return args.Get(0).(internal.Product), args.Error(1)
}

// Save mocks the Create method
func (m *ProductMock) Save(product internal.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

// Update mocks the Update method
func (m *ProductMock) Update(product internal.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

// Delete mocks the Delete method
func (m *ProductMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

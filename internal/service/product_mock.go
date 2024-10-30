package service

import (
	"github.com/manuelfirman/go-API/internal"
	"github.com/stretchr/testify/mock"
)

// ProductMock is a mock type for the ProductRepository interface
type ProductMock struct {
	// Embedding the mock.Mock type allows us to call the Called method
	mock.Mock
}

// GetAll mocks the GetAll method
func (m *ProductMock) GetAll() ([]internal.Product, error) {
	args := m.Called()
	return args.Get(0).([]internal.Product), args.Error(1)
}

// Get mocks the Get method
func (m *ProductMock) Get(id int) (internal.Product, error) {
	args := m.Called(id)
	return args.Get(0).(internal.Product), args.Error(1)
}

// Save mocks the Save method
func (m *ProductMock) Save(product *internal.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

// Update mocks the Update method
func (m *ProductMock) Update(product *internal.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

// Delete mocks the Delete method
func (m *ProductMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

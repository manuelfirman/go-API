package repository

import (
	"github.com/manuelfirman/go-API/internal"
	"github.com/stretchr/testify/mock"
)

// WarehouseMock is a mock for WarehouseRepository
type WarehouseMock struct {
	// Embed the interface to ensure that the mock satisfies the interface
	mock.Mock
}

// GetAll mocks the GetAll method
func (m *WarehouseMock) GetAll() ([]internal.Warehouse, error) {
	// This will return whatever we pass to it
	args := m.Called()
	return args.Get(0).([]internal.Warehouse), args.Error(1)
}

// Get mocks the Get method
func (m *WarehouseMock) Get(id int) (internal.Warehouse, error) {
	args := m.Called(id)
	return args.Get(0).(internal.Warehouse), args.Error(1)
}

// Save mocks the Save method
func (m *WarehouseMock) Save(warehouse *internal.Warehouse) error {
	args := m.Called(warehouse)
	return args.Error(0)
}

// Update mocks the Update method
func (m *WarehouseMock) Update(warehouse *internal.Warehouse) error {
	args := m.Called(warehouse)
	return args.Error(0)
}

// Delete mocks the Delete method
func (m *WarehouseMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

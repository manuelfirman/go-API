package repository

import (
	"github.com/manuelfirman/go-API/internal"
	"github.com/stretchr/testify/mock"
)

// SellerMock is a mock for SellerRepository
type SellerMock struct {
	// Embed the interface to ensure that the mock satisfies the interface
	mock.Mock
}

// GetAll mocks the GetAll method
func (m *SellerMock) GetAll() ([]internal.Seller, error) {
	// This will return whatever we pass to it
	args := m.Called()
	return args.Get(0).([]internal.Seller), args.Error(1)
}

// Get mocks the Get method
func (m *SellerMock) Get(id int) (internal.Seller, error) {
	args := m.Called(id)
	return args.Get(0).(internal.Seller), args.Error(1)
}

// Save mocks the Save method
func (m *SellerMock) Save(seller *internal.Seller) error {
	args := m.Called(seller)
	return args.Error(0)
}

// Update mocks the Update method
func (m *SellerMock) Update(seller *internal.Seller) error {
	args := m.Called(seller)
	return args.Error(0)
}

// Delete mocks the Delete method
func (m *SellerMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

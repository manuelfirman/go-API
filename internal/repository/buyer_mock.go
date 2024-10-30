package repository

import (
	"github.com/manuelfirman/go-API/internal"
	"github.com/stretchr/testify/mock"
)

// BuyerMock is a mock type for the Buyer repository
type BuyerMock struct {
	// Embed the mock type to the BuyerMock
	mock.Mock
}

// GetAll is a mocked method of the Buyer repository that returns a mocked result
func (m *BuyerMock) GetAll() ([]internal.Buyer, error) {
	// Call the mocked method and return the mocked results
	args := m.Called()
	return args.Get(0).([]internal.Buyer), args.Error(1)
}

// Get is a mocked method of the Buyer repository that returns a mocked result
func (m *BuyerMock) Get(id int) (internal.Buyer, error) {
	// Call the mocked method and return the mocked results
	args := m.Called(id)
	return args.Get(0).(internal.Buyer), args.Error(1)
}

// Save is a mocked method of the Buyer repository that returns a mocked result
func (m *BuyerMock) Save(buyer internal.Buyer) error {
	// Call the mocked method and return the mocked results
	args := m.Called(buyer)
	return args.Error(0)
}

// Update is a mocked method of the Buyer repository that returns a mocked result
func (m *BuyerMock) Update(buyer internal.Buyer) error {
	// Call the mocked method and return the mocked results
	args := m.Called(buyer)
	return args.Error(0)
}

// Delete is a mocked method of the Buyer repository that returns a mocked result
func (m *BuyerMock) Delete(id int) error {
	// Call the mocked method and return the mocked results
	args := m.Called(id)
	return args.Error(0)
}

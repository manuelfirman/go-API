package service

import (
	"github.com/manuelfirman/go-API/internal"
	"github.com/stretchr/testify/mock"
)

// NewBuyerMock creates a new BuyerMock
func NewBuyerMock() *BuyerMock {
	return &BuyerMock{}
}

// BuyerMock is a mock type for the Buyer service
type BuyerMock struct {
	mock.Mock
}

// GetAll is a mocked method of the Buyer service that returns a mocked result
func (m *BuyerMock) GetAll() ([]internal.Buyer, error) {
	// Call the mocked method and return the mocked results
	args := m.Called()
	return args.Get(0).([]internal.Buyer), args.Error(1)
}

// Get is a mocked method of the Buyer service that returns a mocked result
func (m *BuyerMock) Get(id int) (internal.Buyer, error) {
	// Call the mocked method and return the mocked results
	args := m.Called(id)
	return args.Get(0).(internal.Buyer), args.Error(1)
}

// Save is a mocked method of the Buyer service that returns a mocked result
func (m *BuyerMock) Save(buyer *internal.Buyer) error {
	// Call the mocked method and return the mocked results
	args := m.Called(buyer)
	return args.Error(0)
}

// Update is a mocked method of the Buyer service that returns a mocked result
func (m *BuyerMock) Update(buyer *internal.Buyer) error {
	// Call the mocked method and return the mocked results
	args := m.Called(buyer)
	return args.Error(0)
}

// Delete is a mocked method of the Buyer service that returns a mocked result
func (m *BuyerMock) Delete(id int) error {
	// Call the mocked method and return the mocked results
	args := m.Called(id)
	return args.Error(0)
}

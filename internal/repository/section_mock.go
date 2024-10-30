package repository

import (
	"github.com/manuelfirman/go-API/internal"
	"github.com/stretchr/testify/mock"
)

// NewSectionMock is a constructor for the SectionMock type
func NewSectionMock() *SectionMock {
	// Return the address of a new SectionMock struct
	return &SectionMock{}
}

// SectionMock is a mock type for the Section repository
type SectionMock struct {
	// Embed the mock type to the SectionMock
	mock.Mock
}

// GetAll is a mocked method of the Section repository that returns a mocked result
func (m *SectionMock) GetAll() ([]internal.Section, error) {
	// Call the mocked method and return the mocked results
	args := m.Called()
	return args.Get(0).([]internal.Section), args.Error(1)
}

// Get is a mocked method of the Section repository that returns a mocked result
func (m *SectionMock) Get(id int) (internal.Section, error) {
	// Call the mocked method and return the mocked results
	args := m.Called(id)
	return args.Get(0).(internal.Section), args.Error(1)
}

// Save is a mocked method of the Section repository that returns a mocked result
func (m *SectionMock) Save(section *internal.Section) error {
	// Call the mocked method and return the mocked results
	args := m.Called(section)
	return args.Error(0)
}

// Update is a mocked method of the Section repository that returns a mocked result
func (m *SectionMock) Update(section *internal.Section) error {
	// Call the mocked method and return the mocked results
	args := m.Called(section)
	return args.Error(0)
}

// Delete is a mocked method of the Section repository that returns a mocked result
func (m *SectionMock) Delete(id int) error {
	// Call the mocked method and return the mocked results
	args := m.Called(id)
	return args.Error(0)
}

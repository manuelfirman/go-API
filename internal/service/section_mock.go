package service

import (
	"github.com/manuelfirman/go-API/internal"
	"github.com/stretchr/testify/mock"
)

// NewSectionMock is a constructor for the SectionMock type
func NewSectionMock() *SectionMock {
	return &SectionMock{}
}

// SectionMock is a mock type for the Section service
type SectionMock struct {
	mock.Mock
}

// GetAll is a mocked method of the Section service that returns a mocked result
func (m *SectionMock) GetAll() ([]internal.Section, error) {
	args := m.Called()
	return args.Get(0).([]internal.Section), args.Error(1)
}

// Get is a mocked method of the Section service that returns a mocked result
func (m *SectionMock) Get(id int) (internal.Section, error) {
	args := m.Called(id)
	return args.Get(0).(internal.Section), args.Error(1)
}

// Save is a mocked method of the Section service that returns a mocked result
func (m *SectionMock) Save(section *internal.Section) error {
	args := m.Called(section)
	return args.Error(0)
}

// Update is a mocked method of the Section service that returns a mocked result
func (m *SectionMock) Update(section *internal.Section) error {
	args := m.Called(section)
	return args.Error(0)
}

// Delete is a mocked method of the Section service that returns a mocked result
func (m *SectionMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

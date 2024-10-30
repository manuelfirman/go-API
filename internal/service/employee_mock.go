package service

import (
	"github.com/manuelfirman/go-API/internal"
	"github.com/stretchr/testify/mock"
)

// NewEmployeeMock creates a new EmployeeMock
func NewEmployeeMock() *EmployeeMock {
	return &EmployeeMock{}
}

// EmployeeMock is a mock for EmployeeRepository
type EmployeeMock struct {
	// Embed the mock object
	mock.Mock
}

// GetAll is a mock for GetAll method in EmployeeRepository
func (e *EmployeeMock) GetAll() ([]internal.Employee, error) {
	// Call the method
	args := e.Called()
	// Return the mocked values
	return args.Get(0).([]internal.Employee), args.Error(1)
}

// Get is a mock for Get method in EmployeeRepository
func (e *EmployeeMock) Get(id int) (internal.Employee, error) {
	// Call the method
	args := e.Called(id)
	// Return the mocked values
	return args.Get(0).(internal.Employee), args.Error(1)
}

// Save is a mock for Save method in EmployeeRepository
func (e *EmployeeMock) Save(employee *internal.Employee) error {
	// Call the method
	args := e.Called(employee)
	// Return the mocked values
	return args.Error(0)
}

// Update is a mock for Update method in EmployeeRepository
func (e *EmployeeMock) Update(employee *internal.Employee) error {
	// Call the method
	args := e.Called(employee)
	// Return the mocked values
	return args.Error(0)
}

// Delete is a mock for Delete method in EmployeeRepository
func (e *EmployeeMock) Delete(id int) error {
	// Call the method
	args := e.Called(id)
	// Return the mocked values
	return args.Error(0)
}

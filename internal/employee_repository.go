package internal

import "errors"

var (
	// ErrEmployeesRepositoryNotFound is returned when the employees are not found
	ErrEmployeesRepositoryNotFound = errors.New("repository: employees not found")
	// ErrEmployeeRepositoryNotFound is returned when the employee is not found
	ErrEmployeeRepositoryNotFound = errors.New("repository: employee not found")
	// ErrEmployeeRepositoryDuplicated is returned when the employee already exists
	ErrEmployeeRepositoryDuplicated = errors.New("repository: employee already exists")
	// ErrEmployeeRepositoryFieldInvalid is returned when the employee field is invalid
	ErrEmployeeRepositoryInvalidField = errors.New("repository: field invalid")
	// ErrInvalidForeingKey is returned when the foreing key is invalid
	ErrEmployeeRepositoryForeignKey = errors.New("repository: invalid foreing key restriction")
)

// EmployeeRepository is an interface that contains the methods that the employee repository should support
type EmployeeRepository interface {
	// FindAll returns all the employees
	GetAll() ([]Employee, error)
	// FindByID returns the employee with the given ID
	Get(id int) (Employee, error)
	// Save saves the given employee
	Save(employee *Employee) error
	// Update updates the given employee
	Update(employee *Employee) error
	// Delete deletes the employee with the given ID
	Delete(id int) error

	// getReportInboudOrder returns the inbound order report for the given employee id or all orders for all employees
	// GetReportInboudOrders(id int) (iorr []InboundOrderReport, err error)
}

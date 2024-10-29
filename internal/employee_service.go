package internal

import "errors"

var (
	// ErrEmployeeServiceInvalidID is returned when the employee ID is invalid
	ErrEmployeeServiceInvalidID = errors.New("service: invalid id")
	// ErrEmployeeServiceFieldRequired is returned when the employee field is required
	ErrEmployeeServiceFieldRequired = errors.New("service: field required")
	// ErrEmployeeServiceNotNegativeField is returned when the employee field is not negative
	ErrEmployeeServiceNotNegativeField = errors.New("service: field can't be negative")
	//ErrEmployeeServiceInternalError is returned when an internal error occurs
	ErrEmployeeServiceInternalError = errors.New("service: internal error")
)

// EmployeeService is an interface that contains the methods that the employee service should support
type EmployeeService interface {
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

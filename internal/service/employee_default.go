package service

import (
	"fmt"

	"github.com/manuelfirman/go-API/internal"
)

// NewEmployeeDefault creates a new instance of the employee service
func NewEmployeeDefault(rp internal.EmployeeRepository) *EmployeeDefault {
	return &EmployeeDefault{
		rp: rp,
	}
}

// NewEmployeeDefault creates a new instance of the employee service
type EmployeeDefault struct {
	rp internal.EmployeeRepository
}

// GetAll returns all employees. Returns an error if the operation fails.
func (s *EmployeeDefault) GetAll() (employees []internal.Employee, err error) {
	employees, err = s.rp.GetAll()
	if err != nil {
		switch err {
		case internal.ErrEmployeeRepository:
			err = fmt.Errorf("%w: %v", internal.ErrEmployeeServiceInternalError, err)
		default:
			err = fmt.Errorf("%w: %v", internal.ErrEmployeeServiceUnknown, err)
		}

		return
	}

	return
}

// Get returns an employee by ID. Returns an error if the employee is not found.
func (s *EmployeeDefault) Get(id int) (employee internal.Employee, err error) {
	employee, err = s.rp.Get(id)
	if err != nil {
		switch err {
		case internal.ErrEmployeeRepositoryNotFound:
			err = fmt.Errorf("%w: %v", internal.ErrEmployeeServiceNotFound, err)
		case internal.ErrEmployeeRepository:
			err = fmt.Errorf("%w: %v", internal.ErrEmployeeServiceInternalError, err)
		default:
			err = fmt.Errorf("%w: %v", internal.ErrEmployeeServiceUnknown, err)
		}

		return
	}

	return
}

// Save saves the given employee. Returns an error if the operation fails.
func (s *EmployeeDefault) Save(employee *internal.Employee) (err error) {
	// validate employee
	if err = validateEmployee(employee); err != nil {
		return
	}

	err = s.rp.Save(employee)
	if err != nil {
		switch err {
		case internal.ErrEmployeeRepositoryDuplicated:
			err = fmt.Errorf("%w: %v", internal.ErrEmployeeServiceDuplicated, "card number id")
		case internal.ErrEmployeeRepository:
			err = fmt.Errorf("%w: %v", internal.ErrEmployeeServiceInternalError, err)
		default:
			err = fmt.Errorf("%w: %v", internal.ErrEmployeeServiceUnknown, err)
		}

		return
	}

	return
}

// Update updates the given employee. Returns an error if the operation fails.
func (s *EmployeeDefault) Update(employee *internal.Employee) (err error) {
	// validate employee
	if err = validateEmployee(employee); err != nil {
		return
	}

	err = s.rp.Update(employee)
	if err != nil {
		switch err {
		case internal.ErrEmployeeRepositoryNotFound:
			err = fmt.Errorf("%w: %v", internal.ErrEmployeeServiceNotFound, err)
		case internal.ErrEmployeeRepository:
			err = fmt.Errorf("%w: %v", internal.ErrEmployeeServiceInternalError, err)
		default:
			err = fmt.Errorf("%w: %v", internal.ErrEmployeeServiceUnknown, err)
		}

		return
	}

	return
}

// Delete deletes the employee with the given ID. Returns an error if the operation fails.
func (s *EmployeeDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	if err != nil {
		switch err {
		case internal.ErrEmployeeRepositoryNotFound:
			err = fmt.Errorf("%w: %v", internal.ErrEmployeeServiceNotFound, err)
		case internal.ErrEmployeeRepository:
			err = fmt.Errorf("%w: %v", internal.ErrEmployeeServiceInternalError, err)
		default:
			err = fmt.Errorf("%w: %v", internal.ErrEmployeeServiceUnknown, err)
		}

		return
	}

	return
}

// validateEmployee validates the employee fields
func validateEmployee(employee *internal.Employee) (err error) {
	// validate employee
	if employee.FirstName == "" || len(employee.FirstName) < 3 {
		err = fmt.Errorf("%w: %v", internal.ErrEmployeeServiceFieldRequired, "first name")
		return
	}

	if employee.LastName == "" || len(employee.LastName) < 3 {
		err = fmt.Errorf("%w: %v", internal.ErrEmployeeServiceFieldRequired, "last name")
		return
	}

	if employee.WarehouseID < 0 {
		err = fmt.Errorf("%w: %v", internal.ErrEmployeeServiceNotNegativeField, "warehouse id")
		return
	}

	return
}

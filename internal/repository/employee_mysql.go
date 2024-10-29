package repository

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/manuelfirman/go-API/internal"
)

// NewEmployeeMySQL creates a new instance of the employee repository
func NewEmployeeMySQL(db *sql.DB) *EmployeeMySQL {
	return &EmployeeMySQL{
		db: db,
	}
}

// EmployeeMySQL is the default implementation of the employee repository
type EmployeeMySQL struct {
	db *sql.DB
}

// GetAll returns all employees. Returns an error if the operation fails.
func (r *EmployeeMySQL) GetAll() (employees []internal.Employee, err error) {
	// execute the query
	query := "SELECT e.`id`, e.`card_number_id`, e.`first_name`, e.`last_name`, e.`warehouse_id` FROM `employees` AS `e`"
	rows, err := r.db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var employee internal.Employee
		err = rows.Scan(&employee.ID, &employee.CardNumberID, &employee.FirstName, &employee.LastName, &employee.WarehouseID)
		if err != nil {
			return
		}

		employees = append(employees, employee)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Get returns an employee by ID. Returns an error if the employee is not found.
func (r *EmployeeMySQL) Get(id int) (e internal.Employee, err error) {
	// execute the query
	query := "SELECT e.`id`, e.`card_number_id`, e.`first_name`, e.`last_name`, e.`warehouse_id` FROM `employees` AS `e` WHERE e.`id` = ?"
	row := r.db.QueryRow(query, id)
	// scan the row and return the employee
	err = row.Scan(&e.ID, &e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			err = internal.ErrEmployeeRepositoryNotFound
		default:
			err = internal.ErrEmployeeRepository
		}

		return
	}

	return
}

// Save receives an employee and saves it.
func (r *EmployeeMySQL) Save(e *internal.Employee) (err error) {
	// execute the query
	query := "INSERT INTO `employees` (`card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, e.CardNumberID, e.FirstName, e.LastName, e.WarehouseID)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrEmployeeRepositoryDuplicated
			default:
				err = internal.ErrEmployeeRepository
			}
		}

		return
	}

	// get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	e.ID = int(id)
	return
}

// Update receives an employee and updates it. Returns an error if the operation fails.
func (r *EmployeeMySQL) Update(e *internal.Employee) (err error) {
	// execute the query
	query := "UPDATE `employees` SET `card_number_id` = ?, `first_name` = ?, `last_name` = ?, `warehouse_id` = ? WHERE `id` = ?"
	_, err = r.db.Exec(query, e.CardNumberID, e.FirstName, e.LastName, e.WarehouseID, e.ID)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1452:
				err = internal.ErrEmployeeRepositoryForeignKey
			case 1062:
				err = internal.ErrEmployeeRepositoryDuplicated
			default:
				err = internal.ErrEmployeeRepository
			}
		}

		return
	}

	return
}

// Delete receives an employee ID and deletes it. Returns an error if the operation fails.
func (r *EmployeeMySQL) Delete(id int) (err error) {
	// execute the query
	query := "DELETE FROM `employees` WHERE `id` = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1451:
				err = internal.ErrEmployeeRepositoryForeignKey
			default:
				err = internal.ErrEmployeeRepository
			}
		}

		return
	}

	// get the number of rows affected
	rows, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rows == 0 {
		err = internal.ErrEmployeeRepositoryNotFound
	}

	return
}

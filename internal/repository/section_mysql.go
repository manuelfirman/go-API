package repository

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/manuelfirman/go-API/internal"
)

// NewSectionMySQL creates a new instance of the section repository for MySQL
func NewSectionMySQL(db *sql.DB) *SectionMySQL {
	return &SectionMySQL{
		db: db,
	}
}

// SectionMySQL is the default implementation of the section repository for MySQL
type SectionMySQL struct {
	db *sql.DB
}

// GetAll returns all the sections
func (r *SectionMySQL) GetAll() (sections []internal.Section, err error) {
	// execute the query
	query := "SELECT `s.id`, `s.section_number`, `s.current_temperature`, `s.minimum_temperature`, `s.current_capacity`, `s.minimum_capacity`, `s.maximum_capacity`, `s.warehouse_id`, `s.product_type_id` FROM `sections` AS `s`"
	rows, err := r.db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()
	// iterate over the rows
	for rows.Next() {
		var section internal.Section
		err = rows.Scan(&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID)
		if err != nil {
			return
		}

		sections = append(sections, section)
	}
	// check if there was an error during the iteration
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Get returns a section by ID
func (r *SectionMySQL) Get(id int) (section internal.Section, err error) {
	// execute the query
	query := "SELECT `s.id`, `s.section_number`, `s.current_temperature`, `s.minimum_temperature`, `s.current_capacity`, `s.minimum_capacity`, `s.maximum_capacity`, `s.warehouse_id`, `s.product_type_id` FROM `sections` AS `s` WHERE `s.id` = ?"
	row := r.db.QueryRow(query, id)

	// scan the row and return the section
	err = row.Scan(&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			err = internal.ErrSectionRepositoryNotFound
		default:
			err = internal.ErrSectionRepository
		}

		return
	}

	return
}

func (r *SectionMySQL) Save(section *internal.Section) (err error) {
	// execute the query
	query := "INSERT INTO `sections` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query, (*section).SectionNumber, (*section).CurrentTemperature, (*section).MinimumTemperature, (*section).CurrentCapacity, (*section).MinimumCapacity, (*section).MaximumCapacity, (*section).WarehouseID, (*section).ProductTypeID)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrSectionRepositoryDuplicated
			default:
				err = internal.ErrSectionRepository
			}
		}

		return
	}

	// get the ID of the inserted section
	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	// set the ID of the section
	section.ID = int(id)

	return
}

// Update receives a section and updates it
func (r *SectionMySQL) Update(section *internal.Section) (err error) {
	query := "UPDATE `sections` SET `section_number` = ?, `current_temperature` = ?, `minimum_temperature` = ?, `current_capacity` = ?, `minimum_capacity` = ?, `maximum_capacity` = ?, `warehouse_id` = ?, `product_type_id` = ? WHERE `id` = ?"
	_, err = r.db.Exec(query, (*section).SectionNumber, (*section).CurrentTemperature, (*section).MinimumTemperature, (*section).CurrentCapacity, (*section).MinimumCapacity, (*section).MaximumCapacity, (*section).WarehouseID, (*section).ProductTypeID, (*section).ID)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrSectionRepositoryDuplicated
			default:
				err = internal.ErrSectionRepository
			}
		}

		return
	}

	return
}

// Delete receives an ID and deletes the section
func (r *SectionMySQL) Delete(id int) (err error) {
	// execute the query
	query := "DELETE FROM `sections` WHERE `id` = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1451:
				err = internal.ErrSectionRepositoryFK
			default:
				err = internal.ErrSectionRepository
			}
		}

		return
	}
	// check if the section was deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}
	// return an error if the section was not found
	if rowsAffected == 0 {
		err = internal.ErrSectionRepositoryNotFound
		return
	}

	return
}

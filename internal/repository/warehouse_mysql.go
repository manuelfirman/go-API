package repository

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/manuelfirman/go-API/internal"
)

// WarehouseMySQL is the default implementation of the Warehouse repository
type WarehouseMySQL struct {
	// db is the database connection
	db *sql.DB
}

// NewWarehouseMySQL creates a new instance of the Warehouse repository
func NewWarehouseMySQL(db *sql.DB) *WarehouseMySQL {
	return &WarehouseMySQL{
		db: db,
	}
}

// GetAll returns all Warehouses
func (w *WarehouseMySQL) GetAll() (warehouses []internal.Warehouse, err error) {
	query := "SELECT `id`, `warehouse_code`, `address`, `telephone`, `minimum_capacity`, `minimum_temperature`, `locality_id` FROM warehouses"
	rows, err := w.db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var w internal.Warehouse
		err = rows.Scan(&w.ID, &w.WarehouseCode, &w.Address, &w.Telephone, &w.MinimumCapacity, &w.MinimumTemperature, &w.LocalityId)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				err = internal.ErrWarehouseRepositoryNotFound
			default:
				err = internal.ErrWarehouseRepositoryUnknown
			}
			return
		}

		warehouses = append(warehouses, w)
	}

	return
}

// Get returns a Warehouse by ID
func (w *WarehouseMySQL) Get(id int) (wh internal.Warehouse, err error) {
	query := "SELECT `id`, `warehouse_code`, `address`, `telephone`, `minimum_capacity`, `minimum_temperature`, `locality_id` FROM warehouses WHERE id = ?"
	row := w.db.QueryRow(query, id)

	// scan the row and return the product
	err = row.Scan(&wh.ID, &wh.WarehouseCode, &wh.Address, &wh.Telephone, &wh.MinimumCapacity, &wh.MinimumTemperature, &wh.LocalityId)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			err = internal.ErrWarehouseRepositoryNotFound
		default:
			err = internal.ErrWarehouseRepositoryUnknown
		}
		return
	}

	return
}

// Save saves a Warehouse
func (w *WarehouseMySQL) Save(wh *internal.Warehouse) (id int, err error) {
	query := "INSERT INTO Warehouses (warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := w.db.Exec(query, wh.WarehouseCode, wh.Address, wh.Telephone, wh.MinimumCapacity, wh.MinimumTemperature, wh.LocalityId)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrWarehouseRepositoryDuplicated
			case 1452:
				err = internal.ErrWarehouseRepositoryForeignKey
			default:
				err = internal.ErrWarehouseRepositoryUnknown
			}
			return
		}
	}

	var lastID int64
	lastID, err = result.LastInsertId()
	if err != nil {
		err = internal.ErrWarehouseRepositoryUnknown
	}

	id = int(lastID)

	return
}

// Update updates a Warehouse
func (r *WarehouseMySQL) Update(s *internal.Warehouse) (err error) {
	query := "UPDATE Warehouses SET warehouse_code = ?, address = ?, telephone = ?, minimum_capacity = ?, minimum_temperature = ?, locality_id = ? WHERE id = ?"
	result, err := r.db.Exec(query, s.WarehouseCode, s.Address, s.Telephone, s.MinimumCapacity, s.MinimumTemperature, s.LocalityId, s.ID)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1452:
				err = internal.ErrWarehouseRepositoryForeignKey
			default:
				err = internal.ErrWarehouseRepositoryUnknown
			}
			return
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		err = internal.ErrWarehouseRepositoryUnknown
	}

	if rowsAffected == 0 {
		err = internal.ErrWarehouseRepositoryNothingToUpdate
	}

	return
}

// Delete deletes a Warehouse by ID
func (r *WarehouseMySQL) Delete(id int) (err error) {
	query := "DELETE FROM Warehouses WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1451:
				err = internal.ErrWarehouseRepositoryForeignKey
			default:
				err = internal.ErrWarehouseRepositoryUnknown
			}
			return
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			err = internal.ErrWarehouseRepositoryNotFound
		default:
			err = internal.ErrWarehouseRepositoryUnknown
		}
		return
	}

	if rowsAffected == 0 {
		err = internal.ErrWarehouseRepositoryNotFound
	}

	return
}

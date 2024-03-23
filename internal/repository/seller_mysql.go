package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/manuelfirman/go-API/internal"
)

// SellerMySQL is the default implementation of the seller repository
type SellerMySQL struct {
	// db is the database connection
	db *sql.DB
}

// NewSellerMySQL creates a new instance of the seller repository
func NewSellerMySQL(db *sql.DB) *SellerMySQL {
	return &SellerMySQL{
		db: db,
	}
}

// GetAll returns all sellers
func (r *SellerMySQL) GetAll() (sellers []internal.Seller, err error) {
	rows, err := r.db.Query("SELECT * FROM sellers")
	if err != nil {
		return nil, fmt.Errorf("error getting sellers: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var s internal.Seller
		err = rows.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityID)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				err = internal.ErrSellerRepositoryNotFound
			case sql.ErrTxDone:
				err = internal.ErrSellerRepositoryTransaction
			case sql.ErrConnDone:
				err = internal.ErrSellerRepositoryConn
			default:
				err = internal.ErrSellerRepositoryUnknown
			}
			return
		}

		sellers = append(sellers, s)
	}

	return
}

// Get returns a seller by ID
func (r *SellerMySQL) Get(id int) (s internal.Seller, err error) {
	query := "SELECT * FROM sellers WHERE id = ?"
	row := r.db.QueryRow(query, id)

	// scan the row and return the product
	err = row.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			err = internal.ErrSellerRepositoryNotFound
		case errors.Is(err, sql.ErrTxDone):
			err = internal.ErrSellerRepositoryTransaction
		case errors.Is(err, sql.ErrConnDone):
			err = internal.ErrSellerRepositoryConn
		default:
			err = internal.ErrSellerRepositoryUnknown
		}
		return
	}

	return
}

// Save saves a seller
func (r *SellerMySQL) Save(s *internal.Seller) (id int, err error) {
	query := "INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query, s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrSellerRepositoryDuplicated
			case 1452:
				err = internal.ErrSellerRepositoryForeignKey
			default:
				err = internal.ErrSellerRepositoryUnknown
			}
			return
		}
	}

	var lastID int64
	lastID, err = result.LastInsertId()
	if err != nil {
		err = internal.ErrSellerRepositoryUnknown
	}

	id = int(lastID)

	return
}

// Update updates a seller
func (r *SellerMySQL) Update(s *internal.Seller) (err error) {
	query := "UPDATE sellers SET cid = ?, company_name = ?, address = ?, telephone = ?, locality_id = ? WHERE id = ?"
	result, err := r.db.Exec(query, s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID, s.ID)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1452:
				err = internal.ErrSellerRepositoryForeignKey
			default:
				err = internal.ErrSellerRepositoryUnknown
			}
			return
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		err = internal.ErrSellerRepositoryUnknown
	}

	if rowsAffected == 0 {
		err = internal.ErrSellerRepositoryNothingToUpdate
	}

	return
}

// Delete deletes a seller by ID
func (r *SellerMySQL) Delete(id int) (err error) {
	query := "DELETE FROM sellers WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1451:
				err = internal.ErrSellerRepositoryForeignKey
			default:
				err = internal.ErrSellerRepositoryUnknown
			}
			return
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			err = internal.ErrSellerRepositoryNotFound
		case errors.Is(err, sql.ErrTxDone):
			err = internal.ErrSellerRepositoryTransaction
		case errors.Is(err, sql.ErrConnDone):
			err = internal.ErrSellerRepositoryConn
		default:
			err = internal.ErrSellerRepositoryUnknown
		}
		return
	}

	if rowsAffected == 0 {
		err = internal.ErrSellerRepositoryNotFound
	}

	return
}

package repository

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/manuelfirman/go-API/internal"
)

type BuyerMySQL struct {
	db *sql.DB
}

// NewBuyerMySQL creates a new instance of the buyer repository
func NewBuyerMySQL(db *sql.DB) *BuyerMySQL {
	return &BuyerMySQL{
		db: db,
	}
}

// GetAll returns all buyers. Returns an error if the operation fails.
func (r *BuyerMySQL) GetAll() (buyers []internal.Buyer, err error) {
	// execute the query
	query := "SELECT b.`id`, b.`card_number_id`, b.`first_name`, b.`last_name` FROM `buyers` AS `b`"
	rows, err := r.db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var buyer internal.Buyer
		err = rows.Scan(&buyer.ID, &buyer.CardNumberID, &buyer.FirstName, &buyer.LastName)
		if err != nil {
			return
		}

		buyers = append(buyers, buyer)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Get returns a buyer by ID. Returns an error if the buyer is not found.
func (r *BuyerMySQL) Get(id int) (b internal.Buyer, err error) {
	// execute the query
	query := "SELECT b.`id`, b.`card_number_id`, b.`first_name`, b.`last_name` FROM `buyers` AS `b` WHERE b.`id` = ?"
	row := r.db.QueryRow(query, id)
	// scan the row and return the buyer
	err = row.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			err = internal.ErrBuyerRepositoryNotFound
		default:
			err = internal.ErrBuyerRepository
		}

		return
	}

	return
}

// Save receives a buyer and saves it
func (r *BuyerMySQL) Save(b *internal.Buyer) (err error) {
	// execute the query
	query := "INSERT INTO `buyers` (`card_number_id`, `first_name`, `last_name`) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, b.CardNumberID, b.FirstName, b.LastName)

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrBuyerRepositoryDuplicated
			default:
				err = internal.ErrBuyerRepository
			}
		}
		return
	}

	// get the ID of the buyer saved
	var id64 int64
	id64, err = result.LastInsertId()
	if err != nil {
		return
	}

	// set the ID of the buyer
	b.ID = int(id64)

	return
}

// Update receives a buyer and updates it. Returns an error if the buyer is not found.
func (r *BuyerMySQL) Update(b *internal.Buyer) (err error) {
	// execute the query
	query := "UPDATE `buyers` SET `card_number_id` = ?, `first_name` = ?, `last_name` = ? WHERE `id` = ?"
	_, err = r.db.Exec(query, b.CardNumberID, b.FirstName, b.LastName, b.ID)

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrBuyerRepositoryDuplicated
			default:
				err = internal.ErrBuyerRepository
			}
		}

		return
	}

	return
}

// Delete receives a buyer ID and deletes it. Returns an error if the buyer is not found.
func (r *BuyerMySQL) Delete(id int) (err error) {
	// execute the query
	query := "DELETE FROM `buyers` WHERE `id` = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1451:
				err = internal.ErrBuyerRepositoryFK
			default:
				err = internal.ErrBuyerRepository
			}
		}

		return
	}

	// check if the buyer was deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		err = internal.ErrBuyerRepository
		return
	}

	// return an error if the buyer was not deleted
	if rowsAffected == 0 {
		err = internal.ErrBuyerRepositoryNotFound
		return
	}

	return
}

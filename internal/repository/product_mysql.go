package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/manuelfirman/go-API/internal"
)

type repository struct {
	db *sql.DB
}

// NewProductMySQL creates a new instance of the product repository
func NewProductMySQL(db *sql.DB) internal.ProductRepository {
	return &repository{
		db: db,
	}
}

// GetAll returns all products. Returns an error if the operation fails.
func (r *repository) GetAll() (products []internal.Product, err error) {
	// set and execute the query
	query := "SELECT * FROM products;"
	rows, err := r.db.Query(query)
	if err != nil {
		return
	}

	// iterate over the rows and append the products
	for rows.Next() {
		p := internal.Product{}
		_ = rows.Scan(&p.ID, &p.ProductCode, &p.Description, &p.Height, &p.Length, &p.Width, &p.Weight, &p.ExpirationRate, &p.FreezingRate, &p.RecomFreezTemp, &p.ProductTypeID, &p.SellerID)
		products = append(products, p)
	}

	// check for errors
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Get returns a product by ID. Returns an error if the product is not found.
func (r *repository) Get(id int) (p internal.Product, err error) {
	// set and execute the query
	query := "SELECT * FROM products WHERE id=?;"
	row := r.db.QueryRow(query, id)

	// scan the row and return the product
	err = row.Scan(&p.ID, &p.ProductCode, &p.Description, &p.Height, &p.Length, &p.Width, &p.Weight, &p.ExpirationRate, &p.FreezingRate, &p.RecomFreezTemp, &p.ProductTypeID, &p.SellerID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			err = internal.ErrProductRepositoryNotFound
		default:
			err = internal.ErrProductRepositoryUnknown
		}
		return
	}

	return
}

// Save receives a product and saves it. It returns the ID of the product saved.
func (r *repository) Save(p *internal.Product) (err error) {
	// set and prepare the query
	query := "INSERT INTO `products` (`product_code`, `description`, `height`, `length`, `width`, `weight`, `expiration_rate`, `freezing_rate`, `recom_freez_temp`, `product_type_id`, `seller_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query, (*p).ProductCode, (*p).Description, (*p).Height, (*p).Length, (*p).Width, (*p).Weight, (*p).ExpirationRate, (*p).FreezingRate, (*p).RecomFreezTemp, (*p).ProductTypeID, (*p).SellerID)

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				fmt.Println(err)
				err = internal.ErrProductRepositoryDuplicated
			case 1452:
				err = internal.ErrSellerRepositoryNotFound
			default:
				err = internal.ErrProductRepositoryUnknown
			}
		}
		return
	}

	var id64 int64
	id64, err = result.LastInsertId()
	if err != nil {
		return
	}

	(*p).ID = int(id64)

	return
}

// Update receives a product and updates it. Returns an error if the product is not found.
func (r *repository) Update(p *internal.Product) (err error) {
	// execute the query
	query := "UPDATE `products` SET `product_code` = ?, `description` = ?, `height` = ?, `length` = ?, `width` = ?, `weight` = ?, `expiration_rate` = ?, `freezing_rate` = ?, `recom_freez_temp` = ?, `product_type_id` = ?, `seller_id` = ? WHERE `id` = ?"
	result, err := r.db.Exec(query, (*p).ProductCode, (*p).Description, (*p).Height, (*p).Length, (*p).Width, (*p).Weight, (*p).ExpirationRate, (*p).FreezingRate, (*p).RecomFreezTemp, (*p).ProductTypeID, (*p).SellerID, (*p).ID)

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			println(err.Error())
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrProductRepositoryDuplicated
			case 1452:
				err = internal.ErrSellerRepositoryNotFound
			default:
				err = internal.ErrProductRepositoryUnknown
			}
		}
		return
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rows == 0 {
		err = internal.ErrProductRepositoryNothingToUpdate
		return
	}

	return
}

// Delete receives a product ID and deletes it. Returns an error if the product is not found.
func (r *repository) Delete(id int) (err error) {
	// execute the query
	result, err := r.db.Exec("DELETE FROM `products` WHERE `id` = ?", id)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1451:
				err = internal.ErrProductRepositoryForeignKey
			}
		}
		return
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return
	}

	if rows == 0 {
		err = internal.ErrProductRepositoryNotFound
		return
	}

	return
}

// GetRecordsByProductReport returns the product records.
func (r *repository) GetRecordsByProductReport(id int) (products []internal.Product, err error) {
	return
}
